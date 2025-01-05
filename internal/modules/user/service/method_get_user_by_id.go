package user_service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	user_redis "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/redis"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	"github.com/redis/go-redis/v9"

	"github.com/google/uuid"
)

func (userService *userService) GetUserByID(ctx context.Context, id string) (res user_model.GetUserResponse, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return user_model.GetUserResponse{}, util_error.NewBadRequest(err, "id must be uuid")
	}

	userRes, err := user_redis.GetUserData(ctx, userService.redisManager, uid.String())
	if err != nil && err != redis.Nil {
		return user_model.GetUserResponse{}, err
	} else if err == nil {
		return user_model.GetUserResponse{
			ID:    userRes.ID,
			Email: userRes.Email,
			Name:  userRes.Name,
		}, nil
	}

	user, err := userService.userRepository.GetUserByID(ctx, uid)
	if err == sql.ErrNoRows {
		return user_model.GetUserResponse{}, util_error.NewNotFound(err, fmt.Sprintf("User with the id of %s is not found", id))
	}
	if err != nil {
		return user_model.GetUserResponse{}, err
	}

	userRes = user_model.GetUserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}

	err = user_redis.SetUserData(ctx, userService.redisManager, uid.String(), userRes, 24*time.Hour)
	if err != nil {
		return user_model.GetUserResponse{}, err
	}

	return user_model.GetUserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
