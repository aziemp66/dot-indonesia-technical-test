package task_service

import (
	"context"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	task_redis "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/redis"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *taskService) GetTaskByID(ctx context.Context, id string, userID string) (task_model.GetTaskResponse, error) {
	taskRes, err := task_redis.GetTaskData(ctx, s.redisManager, userID, id)
	if err != nil && err != redis.Nil {
		return task_model.GetTaskResponse{}, err
	} else if err == nil {
		if taskRes.UserID != userID {
			return task_model.GetTaskResponse{}, util_error.NewForbidden(err, "You are not the owner of this task")
		}
		return taskRes, nil
	}

	taskID, err := uuid.Parse(id)
	if err != nil {
		return task_model.GetTaskResponse{}, util_error.NewBadRequest(err, "id must be uuid")
	}

	// If not in cache, fetch from the repository
	task, err := s.taskRepository.GetTaskByID(ctx, taskID)
	if err != nil {
		return task_model.GetTaskResponse{}, err
	}

	if task.UserID.String() != userID {
		return task_model.GetTaskResponse{}, util_error.NewForbidden(err, "You are not the owner of this task")
	}

	taskRes = task_model.GetTaskResponse{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      userID,
	}

	// Example: Cache the newly created task (optional)
	// s.redisManager.Set("task:"+id, taskData, expiration)
	err = task_redis.AddTask(ctx, s.redisManager, taskRes)
	if err != nil {
		return task_model.GetTaskResponse{}, err
	}
	// Optionally cache the task after retrieving it
	// s.redisManager.Set("task:"+id, task, expiration)

	return taskRes, nil
}
