package user_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *userRepositoryPostgres) GetUserByID(ctx context.Context, id uuid.UUID) (user_model.User, error) {
	user := user_model.User{
		ID: id,
	}
	if err := r.db.First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return user, util_error.NewNotFound(err, "user not found")
	} else if err != nil {
		return user, fmt.Errorf("error when finding user by ID: %w", err)
	}
	return user, nil
}
