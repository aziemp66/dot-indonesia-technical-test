package task_redis

import (
	"context"
	"encoding/json"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

func GetAllTask(ctx context.Context, r util_redis.RedisManager, userID string) ([]task_model.GetTaskResponse, error) {
	key := getTaskKey(userID)

	fields, err := r.Client().HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var tasks []task_model.GetTaskResponse

	for _, v := range fields {
		var task task_model.GetTaskResponse
		err = json.Unmarshal([]byte(v), &task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
