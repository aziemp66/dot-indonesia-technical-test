package task_service

import (
	"context"

	task_redis "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/redis"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	"github.com/google/uuid"
)

func (s *taskService) DeleteTask(ctx context.Context, id, userID string) error {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return util_error.NewBadRequest(err, "id must be uuid")
	}

	// Call repository to delete the task
	err = s.taskRepository.DeleteTask(ctx, taskID)
	if err != nil {
		return err
	}

	err = task_redis.DeleteTask(ctx, s.redisManager, userID, id)
	if err != nil {
		return err
	}

	err = task_redis.DeleteAllTask(ctx, s.redisManager, userID)
	if err != nil {
		return err
	}

	return nil
}
