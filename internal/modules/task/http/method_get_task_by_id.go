package task_http

import (
	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"
	"github.com/gin-gonic/gin"
)

func (taskHttpHandler *taskHttpHandler) GetTaskByID(ctx *gin.Context) {
	var req task_model.GetTaskIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(util_error.NewBadRequest(err, err.Error()))
		return
	}

	userID := ctx.GetString("user_id")

	taskRes, err := taskHttpHandler.taskService.GetTaskByID(ctx.Request.Context(), req.ID, userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully retrieved Task", taskRes)
}
