package user_http

import (
	user_service "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/service"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
)

func NewUserHttpHandler(userService user_service.UserService, jwtManager util_jwt.JWTManager) UserHttpHandler {
	return &userHttpHandler{
		userService: userService,
		jwtManager:  jwtManager,
	}
}
