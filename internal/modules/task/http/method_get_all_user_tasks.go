package task_http

import (
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"
	"github.com/gin-gonic/gin"
)

func (taskHttpHandler *taskHttpHandler) GetAllUserTasks(ctx *gin.Context) {
	userID := ctx.GetString("user_id")

	tasksRes, err := taskHttpHandler.taskService.GetAllUserTasks(ctx.Request.Context(), userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully retrieved all user tasks", tasksRes)
}
