package telegramHelper

import (
	"api/common/helpers"
	"context"
	"encoding/json"
	"fmt"
	"time"

	goredislib "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type TelegramJob struct {
	BotToken   string
	ChatID     int64
	Message    string
	Attempts   int
	MaxAttempt int
}

func (job *TelegramJob) serialize() ([]byte, error) {
	return json.Marshal(job)
}

func deserialize(data []byte) (*TelegramJob, error) {
	var job TelegramJob
	err := json.Unmarshal(data, &job)
	return &job, err
}

func safeStartWorker(redisClient *goredislib.Client, botToken string) {
	lockKey := fmt.Sprintf("telegram:lock:%s", botToken)
	lockValue := fmt.Sprintf("worker-%d", time.Now().UnixNano()) // unique lock ID
	lockTTL := 10 * time.Second
	refreshInterval := 5 * time.Second

	ctx := context.Background()

	// Attempt to acquire lock
	acquired, err := redisClient.SetNX(ctx, lockKey, lockValue, lockTTL).Result()
	if err != nil {
		helpers.Logger.Error("Telegram error acquiring lock:", zap.Error(err))
		return
	}
	if !acquired {
		helpers.Logger.Error("Telegram worker already running for token:", zap.String("Token", botToken))
		return
	}

	helpers.Logger.Info("Telegram lock acquired. Starting worker.")

	// Start a goroutine to auto-refresh the lock
	stopRenewal := make(chan struct{})
	go func() {
		ticker := time.NewTicker(refreshInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Renew the lock if still owned
				val, err := redisClient.Get(ctx, lockKey).Result()
				if err == nil && val == lockValue {
					redisClient.Expire(ctx, lockKey, lockTTL)
				}
			case <-stopRenewal:
				return
			}
		}
	}()

	// Start the actual blocking worker
	startRetryWorker(redisClient, botToken)

	// When worker ends (never normally), release lock
	close(stopRenewal)
	val, err := redisClient.Get(ctx, lockKey).Result()
	if err == nil && val == lockValue {
		redisClient.Del(ctx, lockKey)
	}
}

func startRetryWorker(redisClient *goredislib.Client, botToken string) {
	retryQueue := "telegram:queue:retry:" + botToken
	ctx := context.Background()

	for {
		data, err := redisClient.BLPop(ctx, 0, retryQueue).Result()
		if err != nil {
			helpers.Logger.Error("Telegram failed to pop from retry queue", zap.Error(err))
			continue
		}

		jobData := []byte(data[1])
		job, err := deserialize(jobData)
		if err != nil {
			helpers.Logger.Error("Telegram failed to parse retry job", zap.Error(err))
			continue
		}

		_, err = postHttpMessage(job)
		if err != nil {
			job.Attempts++
			if job.Attempts >= job.MaxAttempt {
				helpers.Logger.Error("Telegram retry failed - max attempts exceeded", zap.Int64("Chat ID", job.ChatID))
				// Optionally: push to dead-letter queue
				continue
			}
			// Requeue again
			newData, _ := job.serialize()
			_ = redisClient.RPush(ctx, retryQueue, newData).Err()
		} else {
			helpers.Logger.Info("Telegram retry succeeded", zap.Int64("Chat ID", job.ChatID))
		}

		// Always respect Telegram rate limit
		time.Sleep(2 * time.Second)
	}
}

func pushJob(redisClient *goredislib.Client, job *TelegramJob) error {
	ctx := context.Background()
	data, err := job.serialize()
	if err != nil {
		return err
	}
	// We queue the job in a Redis list named per-token
	queueName := "telegram:queue:" + job.BotToken
	return redisClient.RPush(ctx, queueName, data).Err()
}
