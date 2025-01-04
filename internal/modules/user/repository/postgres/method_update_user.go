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

func (r *userRepositoryPostgres) UpdateUser(ctx context.Context, id uuid.UUID, name string) error {
	user := &user_model.User{
		Name: name,
	}
	if err := r.db.WithContext(ctx).Model(&user_model.User{}).Where("id = ?", id).Updates(user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return util_error.NewNotFound(err, "user not found")
	} else if err != nil {
		return fmt.Errorf("error when updating user: %w", err)
	}
	return nil
}
