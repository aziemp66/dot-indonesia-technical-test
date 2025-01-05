package task_redis

import (
	"context"
	"encoding/json"
	"time"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func SetAllTask(ctx context.Context, r util_redis.RedisManager, userID string, task []task_model.GetTaskResponse, expiration time.Duration) error {
	key := getAllTaskKey(userID)

	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return r.Client().Set(ctx, key, data, expiration).Err()
}
