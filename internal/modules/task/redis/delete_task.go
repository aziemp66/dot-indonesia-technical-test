package task_redis

import (
	"context"

	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func DeleteTask(ctx context.Context, r util_redis.RedisManager, userID string, taskID string) error {
	key := getTaskKey(userID)
	return r.Client().HDel(ctx, key, taskID).Err()
}
