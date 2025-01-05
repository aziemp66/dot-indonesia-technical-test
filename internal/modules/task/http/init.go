package task_http

import (
	task_service "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/service"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
)

func NewTaskHttpHandler(taskService task_service.TaskService, jwtManager util_jwt.JWTManager) TaskHttpHandler {
	return &taskHttpHandler{
		taskService: taskService,
		jwtManager:  jwtManager,
	}
}
