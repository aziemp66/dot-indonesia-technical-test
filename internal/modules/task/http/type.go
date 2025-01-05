package task_http

import (
	task_service "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/service"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
	"github.com/gin-gonic/gin"
)

type taskHttpHandler struct {
	taskService task_service.TaskService
	jwtManager  util_jwt.JWTManager
}

type TaskHttpHandler interface {
	CreateTask(ctx *gin.Context)
	GetTaskByID(ctx *gin.Context)
	GetAllUserTasks(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	Deletetask(ctx *gin.Context)
}
