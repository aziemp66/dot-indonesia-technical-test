package user_service

import (
	user_repository "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/repository"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
	util_password "github.com/aziemp66/dot-indonesia-technical-test/util/password"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

type userService struct {
	userRepository  user_repository.UserRepository
	jwtManager      util_jwt.JWTManager
	passwordManager util_password.PasswordManager
	redisManager    util_redis.RedisManager
}

func NewUserService(userRepository user_repository.UserRepository, jwtManager util_jwt.JWTManager, passwordManager util_password.PasswordManager, redisManager util_redis.RedisManager) UserService {
	return &userService{
		userRepository:  userRepository,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
		redisManager:    redisManager,
	}
}
