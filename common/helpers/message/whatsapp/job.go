package whatsappHelper

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"api/common/helpers"

	goredislib "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type WhatsAppJob struct {
	PhoneID             string
	AccessToken         string
	ReceiverPhoneNumber string

	Template   string
	Language   string
	BodyParams []string
	TTLSeconds int

	Attempts   int
	MaxAttempt int
}

func (job *WhatsAppJob) serialize() ([]byte, error) {
	return json.Marshal(job)
}

func deserialize(data []byte) (*WhatsAppJob, error) {
	var job WhatsAppJob
	err := json.Unmarshal(data, &job)
	return &job, err
}

func safeStartWorker(redisClient *goredislib.Client, accessToken, phoneID string) {
	lockKey := fmt.Sprintf("whatsapp:lock:%s", phoneID)
	lockValue := fmt.Sprintf("worker-%d", time.Now().UnixNano())
	lockTTL := 10 * time.Second
	refreshInterval := 5 * time.Second

	ctx := context.Background()
	acquired, err := redisClient.SetNX(ctx, lockKey, lockValue, lockTTL).Result()
	if err != nil {
		helpers.Logger.Error("WhatsApp error acquiring lock:", zap.Error(err))
		return
	}
	if !acquired {
		helpers.Logger.Error("WhatsApp worker already running for phoneID:", zap.String("PhoneID", phoneID))
		return
	}

	helpers.Logger.Info("WhatsApp lock acquired. Starting worker.")

	stopRenewal := make(chan struct{})
	go func() {
		ticker := time.NewTicker(refreshInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				val, err := redisClient.Get(ctx, lockKey).Result()
				if err == nil && val == lockValue {
					redisClient.Expire(ctx, lockKey, lockTTL)
				}
			case <-stopRenewal:
				return
			}
		}
	}()

	startRetryWorker(redisClient, accessToken, phoneID)

	close(stopRenewal)
	val, err := redisClient.Get(ctx, lockKey).Result()
	if err == nil && val == lockValue {
		redisClient.Del(ctx, lockKey)
	}
}

func startRetryWorker(redisClient *goredislib.Client, accessToken, phoneID string) {
	queueName := "whatsapp:queue:" + phoneID
	ctx := context.Background()

	for {
		data, err := redisClient.BLPop(ctx, 0, queueName).Result()
		if err != nil {
			helpers.Logger.Error("WhatsApp failed to pop from queue", zap.Error(err))
			continue
		}

		jobData := []byte(data[1])
		job, err := deserialize(jobData)
		if err != nil {
			helpers.Logger.Error("WhatsApp failed to parse job", zap.Error(err))
			continue
		}

		respData, errPost := postTemplateMessage(job)
		if errPost != nil {
			helpers.Logger.Error("WhatsApp failed to post message", zap.Error(errPost))
			job.Attempts++
			if job.Attempts >= job.MaxAttempt {
				helpers.Logger.Error("WhatsApp failed - max attempts exceeded", zap.String("Phone Number", job.ReceiverPhoneNumber))
				// Optionally: push to dead-letter queue
				continue
			}
			// Requeue again
			newData, _ := job.serialize()
			_ = redisClient.RPush(ctx, queueName, newData).Err()
		} else {
			helpers.Logger.Info(
				fmt.Sprintf("WhatsApp send succeeded after attempt %d", job.Attempts),
				zap.String("Phone Number", job.ReceiverPhoneNumber),
				zap.Any("Response data", respData),
			)
		}

		// Always respect rate limit: ~1 msg/sec
		time.Sleep(2 * time.Second) // Sleep every 2sec
	}
}

func pushJob(redisClient *goredislib.Client, job *WhatsAppJob) error {
	ctx := context.Background()
	data, err := job.serialize()
	if err != nil {
		return err
	}
	queueName := "whatsapp:queue:" + job.PhoneID
	return redisClient.RPush(ctx, queueName, data).Err()
}
