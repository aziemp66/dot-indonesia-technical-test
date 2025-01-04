package user_service

import (
	user_repository "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/repository"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
	util_password "github.com/aziemp66/dot-indonesia-technical-test/util/password"
)

type userService struct {
	userRepository  user_repository.UserRepository
	jwtManager      util_jwt.JWTManager
	passwordManager util_password.PasswordManager
}

func NewUserService(userRepository user_repository.UserRepository, jwtManager util_jwt.JWTManager, passwordManager util_password.PasswordManager) UserService {
	return &userService{
		userRepository:  userRepository,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}
