package task_service

import (
	"context"

	task_redis "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/redis"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	"github.com/google/uuid"
)

func (s *taskService) CreateTask(ctx context.Context, userID string, title, description, status string) (id string, err error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return "", util_error.NewBadRequest(err, "id must be uuid")
	}

	// Call repository to create the task
	id, err = s.taskRepository.CreateTask(ctx, uid, title, description, status)
	if err != nil {
		return "", err
	}

	err = task_redis.DeleteAllTask(ctx, s.redisManager, userID)
	if err != nil {
		return "", err
	}

	return id, nil
}
