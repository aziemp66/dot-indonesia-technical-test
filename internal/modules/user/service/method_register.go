package user_service

import (
	"context"
	"database/sql"
	"fmt"

	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
)

func (userService *userService) Register(ctx context.Context, email, password, name string) (id string, err error) {
	_, err = userService.userRepository.GetUserByEmail(ctx, email)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	} else if err == nil {
		return "", util_error.NewBadRequest(fmt.Errorf("%s is already registered", email), fmt.Sprintf("Email %s is already used", email))
	}

	if err := userService.passwordManager.PasswordValidation(password); err != nil {
		return "", err
	}

	hashedPassword, err := userService.passwordManager.HashPassword(password)
	if err != nil {
		return "", err
	}

	id, err = userService.userRepository.CreateUser(ctx, email, hashedPassword, name)
	if err != nil {
		return "", err
	}

	return id, nil
}
