package task_redis

import (
	"context"
	"encoding/json"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func GetTaskData(ctx context.Context, r util_redis.RedisManager, userID string, taskID string) (task_model.GetTaskResponse, error) {
	key := getTaskKey(userID)

	data, err := r.Client().HGet(ctx, key, taskID).Result()
	if err != nil {
		return task_model.GetTaskResponse{}, nil
	}

	var task task_model.GetTaskResponse

	err = json.Unmarshal([]byte(data), &task)
	if err != nil {
		return task_model.GetTaskResponse{}, nil
	}

	return task, nil
}
