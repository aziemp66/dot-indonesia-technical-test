package user_service

import (
	"context"

	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"github.com/google/uuid"
)

func (userService *userService) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) (err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return util_error.NewBadRequest(err, "id must be uuid")
	}

	user, err := userService.userRepository.GetUserByID(ctx, uid)
	if err != nil {
		return err
	}

	if err := userService.passwordManager.CheckPasswordHash(oldPassword, user.Password); err != nil {
		return err
	}

	if err := userService.passwordManager.PasswordValidation(newPassword); err != nil {
		return err
	}

	hashedPassword, err := userService.passwordManager.HashPassword(newPassword)
	if err != nil {
		return err
	}

	err = userService.userRepository.ChangePassword(ctx, uid, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
