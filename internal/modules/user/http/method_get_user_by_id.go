package user_http

import (
	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) GetUserByID(ctx *gin.Context) {
	var req user_model.GetUserIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(util_error.NewBadRequest(err, err.Error()))
		return
	}

	user, err := userHttpHandler.userService.GetUserByID(ctx.Request.Context(), req.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Retrieve User", user)
}
