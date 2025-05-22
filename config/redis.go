package config

import (
	"api/common/helpers"
	"context"
	"fmt"
	"slices"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

// Establishes a connection to the Redis server.
func ConnectRedis() error {
	addr := fmt.Sprintf("%s:%d", Env.RedisHost, Env.RedisPort)
	helpers.Logger.Warn(
		"Redis address: ",
		zap.String("Value: ", addr),
	)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: Env.RedisUsername,
		Password: Env.RedisPassword,
		DB:       Env.RedisDatabase,
	})

	// Check redis status
	err := CheckRedis()
	return err
}

// Check redis status
func CheckRedis() (err error) {
	err = RedisClient.Ping(RedisContext).Err()
	return
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
	return RedisClient.Get(RedisContext, key).Result()
}

// Stores a string value in Redis using the provided key.
func SetRedisString(key string, val string) error {
	return RedisClient.Set(RedisContext, key, val, 0).Err()
}

// Removes a string value from Redis using the provided key.
func DeleteRedisString(key string) (int64, error) {
	return RedisClient.Del(RedisContext, key).Result()
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
	len, errLen := RedisClient.LLen(RedisContext, key).Result()
	if errLen != nil {
		return nil, errLen
	}
	return RedisClient.LRange(RedisContext, key, 0, len-1).Result()
}

// Appends a new element to a string array stored in Redis.
func AppendToRedisStringList(key string, val string) error {
	return RedisClient.LPush(RedisContext, key, val).Err()
}

// Removes the element at the specified index from a string array stored in Redis.
func RemoveFromRedisStringList(key string, index int64) error {
	_, err := RedisClient.LTrim(RedisContext, key, index, index).Result()
	return err
}
