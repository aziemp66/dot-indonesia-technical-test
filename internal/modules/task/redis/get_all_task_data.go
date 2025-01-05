package task_redis

import (
	"context"
	"encoding/json"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func GetAllTask(ctx context.Context, r util_redis.RedisManager, userID string) ([]task_model.GetTaskResponse, error) {
	key := getAllTaskKey(userID)

	data, err := r.Client().Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var tasks []task_model.GetTaskResponse

	err = json.Unmarshal([]byte(data), &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
