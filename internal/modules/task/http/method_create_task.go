package task_http

import (
	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"
	"github.com/gin-gonic/gin"
)

func (taskHttpHandler *taskHttpHandler) CreateTask(ctx *gin.Context) {
	var req task_model.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	userID := ctx.GetString("user_id")

	id, err := taskHttpHandler.taskService.CreateTask(ctx.Request.Context(), userID, req.Title, req.Description, req.Title)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Create task", map[string]string{
		"task_id": id,
	})
}
