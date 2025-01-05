package task_service

import (
	"context"
	"time"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	task_redis "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/redis"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (s *taskService) GetAllUserTasks(ctx context.Context, userID string) ([]task_model.GetTaskResponse, error) {
	tasksRes, err := task_redis.GetAllTask(ctx, s.redisManager, userID)
	if err != nil && err != redis.Nil {
		return nil, err
	} else if err == nil {
		return tasksRes, nil
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, util_error.NewBadRequest(err, "id must be uuid")
	}
	// Cache could be used for storing frequent user task data
	// cachedTasks, err := s.redisManager.Get("user_tasks:" + userID)
	// if err == nil {
	//     return cachedTasks, nil
	// }

	// Fetch from repository if not in cache
	tasks, err := s.taskRepository.GetAllUserTasks(ctx, uid)
	if err != nil {
		return nil, err
	}

	for _, v := range tasks {
		tasksRes = append(tasksRes, task_model.GetTaskResponse{
			ID:          v.ID.String(),
			Title:       v.Title,
			Description: v.Description,
			Status:      v.Status,
			UserID:      v.UserID.String(),
		})
	}

	err = task_redis.SetAllTask(ctx, s.redisManager, userID, tasksRes, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return tasksRes, nil
}
