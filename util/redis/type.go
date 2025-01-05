package util_redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisManager interface {
	Client() *redis.Client
	Close() error
}
