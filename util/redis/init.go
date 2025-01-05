package util_redis

import (
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

func (r *redisManager) Client() *redis.Client {
	return r.client
}

func (r *redisManager) Close() error {
	return r.client.Close()
}
