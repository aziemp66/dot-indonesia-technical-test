package user_http

import (
	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) Login(ctx *gin.Context) {
	var req user_model.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	token, err := userHttpHandler.userService.Login(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Sucessfully Logged In", user_model.LoginUserResponse{
		Token: token,
	})
}
