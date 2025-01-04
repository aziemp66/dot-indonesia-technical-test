package user_service

import (
	"context"
	"database/sql"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"github.com/google/uuid"
)

func (userService *userService) GetUserByID(ctx context.Context, id string) (res user_model.GetUserResponse, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return user_model.GetUserResponse{}, util_error.NewBadRequest(err, "id must be uuid")
	}

	user, err := userService.userRepository.GetUserByID(ctx, uid)
	if err == sql.ErrNoRows {
		return user_model.GetUserResponse{}, util_error.NewNotFound(err, fmt.Sprintf("User with the id of %s is not found", id))
	}
	if err != nil {
		return user_model.GetUserResponse{}, err
	}

	return user_model.GetUserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
