package user_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"gorm.io/gorm"
)

func (r *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (user_model.User, error) {
	user := user_model.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return user, util_error.NewNotFound(err, "user not found")
	} else if err != nil {
		return user, fmt.Errorf("error when finding user by email: %w", err)
	}
	return user, nil
}
