package task_redis

import (
	"context"

	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func DeleteAllTask(ctx context.Context, r util_redis.RedisManager, userID string) error {
	key := getAllTaskKey(userID)
	return r.Client().Del(ctx, key, key).Err()
}
