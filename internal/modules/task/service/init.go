package task_service

import (
	task_repository "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/repository"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

type taskService struct {
	taskRepository task_repository.TaskRepository
	redisManager   util_redis.RedisManager
}

func NewTaskService(taskRepository task_repository.TaskRepository, redisManager util_redis.RedisManager) TaskService {
	return &taskService{
		taskRepository: taskRepository,
		redisManager:   redisManager,
	}
}
