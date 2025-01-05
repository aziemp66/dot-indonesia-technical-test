package task_http_router

import (
	task_http "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/http"
	util_http_middleware "github.com/aziemp66/dot-indonesia-technical-test/util/http/middleware"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
	"github.com/gin-gonic/gin"
)

func BindTaskHttpRouter(router *gin.RouterGroup, taskHttpHandler task_http.TaskHttpHandler, jwtManager util_jwt.JWTManager) {
	authRouter := router.Group(
		"",
		util_http_middleware.JWTAuthentication(jwtManager),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)

	authRouter.GET("", taskHttpHandler.GetAllUserTasks)
	authRouter.GET("/:id", taskHttpHandler.GetTaskByID)
	authRouter.POST("", taskHttpHandler.CreateTask)
	authRouter.PUT("", taskHttpHandler.UpdateTask)
	authRouter.DELETE("", taskHttpHandler.Deletetask)
}
