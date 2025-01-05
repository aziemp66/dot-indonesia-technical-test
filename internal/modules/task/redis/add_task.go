package task_redis

import (
	"context"
	"encoding/json"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func AddTask(ctx context.Context, r util_redis.RedisManager, task task_model.GetTaskResponse) error {
	key := getTaskKey(task.UserID)

	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return r.Client().HSet(ctx, key, task.ID, data).Err()
}
