package user_service

import (
	"context"
	"database/sql"
	"fmt"

	user_redis "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/redis"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"github.com/google/uuid"
)

func (userService *userService) UpdateUser(ctx context.Context, id, name string) (err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return util_error.NewBadRequest(err, "id must be uuid")
	}

	_, err = userService.userRepository.GetUserByID(ctx, uid)
	if err == sql.ErrNoRows {
		return util_error.NewNotFound(fmt.Errorf("user with the id of %s is not found", id), "User not found")
	} else if err != nil {
		return err
	}

	err = userService.userRepository.UpdateUser(ctx, uid, name)
	if err != nil {
		return err
	}

	err = user_redis.DeleteUserData(ctx, userService.redisManager, id)
	if err != nil {
		return err
	}

	return nil
}
