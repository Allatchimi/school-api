package config

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

var RedisClient *goredislib.Client
var redsyncInstance *redsync.Redsync

const (
	gitDeployLockKey = "git-deploy-lock"
)

// Establishes a connection to the Redis server.
func ConnectRedis() error {
	addr := fmt.Sprintf("%s:%d", Env.RedisHost, Env.RedisPort)
	RedisClient = goredislib.NewClient(&goredislib.Options{
		Addr:     addr,
		Username: Env.RedisUsername,
		Password: Env.RedisPassword,
		DB:       Env.RedisDatabase,
	})

	pool := goredis.NewPool(RedisClient)
	redsyncInstance = redsync.New(pool)

	// Check redis status
	err := CheckRedis()
	return err
}

// Check redis status
func CheckRedis() (err error) {
	err = RedisClient.Ping(context.Background()).Err()
	return
}

// GitDistributedLock acquires a distributed lock to prevent multiple git operations at the same time.
func GitDistributedLock(operation func() error) error {
	mutex := redsyncInstance.NewMutex(gitDeployLockKey,
		redsync.WithExpiry(30*time.Second),
		redsync.WithTries(5),
	)

	if err := mutex.Lock(); err != nil {
		errMsg := "Failed to acquire lock!"
		return fmt.Errorf("%s: %s %w", errMsg, err.Error(), err)
	}
	defer mutex.Unlock()

	return operation()
}

//
//
//
//
// String operations below
//
//
//
//

// Retrieves a string value from Redis using the provided key.
func GetRedisString(key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}

// Stores a string value in Redis using the provided key.
func SetRedisString(key string, val string) error {
	return RedisClient.Set(context.Background(), key, val, 0).Err()
}

// Removes a string value from Redis using the provided key.
func DeleteRedisString(key string) (int64, error) {
	return RedisClient.Del(context.Background(), key).Result()
}

//
//
//
//
// String List operations below
//
//
//

// Creates a function that checks if a specified value exists in a Redis list.
//
// The returned function takes a Redis key as input and returns the found value
// (or an empty string if not found) and an error if any.
func CheckValueInRedisList(requiredVal string) func(string) (string, error) {
	return func(key string) (string, error) {
		array, errArray := GetRedisStringList(key)
		if errArray != nil {
			return "", errArray
		}
		if slices.Contains(array, requiredVal) {
			return requiredVal, nil
		}
		return "", nil
	}
}

// Retrieves a string array from Redis using the provided key.
func GetRedisStringList(key string) ([]string, error) {
	len, errLen := RedisClient.LLen(context.Background(), key).Result()
	if errLen != nil {
		return nil, errLen
	}
	return RedisClient.LRange(context.Background(), key, 0, len-1).Result()
}

// Appends a new element to a string array stored in Redis.
func AppendToRedisStringList(key string, val string) error {
	return RedisClient.LPush(context.Background(), key, val).Err()
}

// Removes the element at the specified index from a string array stored in Redis.
func RemoveFromRedisStringList(key string, index int64) error {
	_, err := RedisClient.LTrim(context.Background(), key, index, index).Result()
	return err
}
