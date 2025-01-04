package user_http

import (
	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) UpdateProfile(ctx *gin.Context) {
	var req user_model.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	userID := ctx.GetString("user_id")

	err := userHttpHandler.userService.UpdateUser(ctx.Request.Context(), userID, req.Name)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Updated Profile", nil)
}
