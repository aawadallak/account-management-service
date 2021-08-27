package redis

import (
	"context"
	"encoding/json"
	"errors"
	"latest/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisConfig struct {
	client *redis.Client
}

func RedisClient() *redisConfig {
	return &redisConfig{
		client: redis.NewClient(&redis.Options{
			Addr:     config.GetConfig().RedisAdress,
			Password: config.GetConfig().RedisPassword,
			DB:       int(config.GetConfig().RedisDatabase),
		})}
}

func (r *redisConfig) Writer(key string, value interface{}, expirationTime time.Duration) error {

	ctx := context.Background()

	t, err := json.Marshal(value)

	if err != nil {
		return err
	}

	err = r.client.Set(ctx, key, t, expirationTime).Err()
	if err != nil {
	}

	return nil
}

func (r *redisConfig) Find(key string) (string, error) {

	ctx := context.Background()

	t, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", errors.New("Invalid code. Please, try again.")
	}

	if err != nil {
		return "", err
	}

	return t, nil

}
