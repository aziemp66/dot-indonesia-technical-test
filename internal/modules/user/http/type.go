package user_http

import (
	user_service "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/service"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"

	"github.com/gin-gonic/gin"
)

type userHttpHandler struct {
	userService user_service.UserService
	jwtManager  util_jwt.JWTManager
}

type UserHttpHandler interface {
	ChangePassword(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}
