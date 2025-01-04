package util_redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisManager struct {
	client *redis.Client
}

func NewRedisManager(addr string, password string, db int) RedisManager {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisManager{
		client: client,
	}
}

func (r *redisManager) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// Serialize value to JSON
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.client.Set(ctx, key, data, expiration).Err()
	return err
}

func (r *redisManager) Get(ctx context.Context, key string, dest interface{}) error {
	// Get JSON string from Redis
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil // Key does not exist
	}
	if err != nil {
		return err
	}

	// Deserialize JSON to the provided destination object
	err = json.Unmarshal([]byte(data), dest)
	return err
}

func (r *redisManager) Delete(ctx context.Context, key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}

func (r *redisManager) Close() error {
	return r.client.Close()
}
