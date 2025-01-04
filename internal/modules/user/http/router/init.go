package user_http_router

import (
	user_http "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/http"
	util_http_middleware "github.com/aziemp66/dot-indonesia-technical-test/util/http/middleware"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"

	"github.com/gin-gonic/gin"
)

func BindUserHttpRouter(router *gin.RouterGroup, handler user_http.UserHttpHandler, jwtManager util_jwt.JWTManager) {
	router.GET("/:id", handler.GetUserByID)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	userRoutes := router.Group(
		"",
		util_http_middleware.JWTAuthentication(jwtManager),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)
	userRoutes.PUT("", handler.UpdateProfile)
	userRoutes.POST("/change-password", handler.ChangePassword)
}
