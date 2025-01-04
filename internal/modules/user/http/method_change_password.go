package user_http

import (
	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) ChangePassword(ctx *gin.Context) {
	id := ctx.GetString("user_id")

	var req user_model.ChangePasswordRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	err := userHttpHandler.userService.ChangePassword(ctx.Request.Context(), id, req.OldPassword, req.NewPassword)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Change Password", nil)
}
