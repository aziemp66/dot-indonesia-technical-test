package task_http

import (
	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"
	"github.com/gin-gonic/gin"
)

func (taskHttpHandler *taskHttpHandler) UpdateTask(ctx *gin.Context) {
	var req task_model.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	userID := ctx.GetString("user_id")

	err := taskHttpHandler.taskService.UpdateTask(ctx.Request.Context(), req.ID, userID, req.Title, req.Description, req.Status)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully updated task", nil)
}
