package user_repository_postgres

import (
	"context"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
)

func (r *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (user_model.User, error) {
	user := user_model.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return user, fmt.Errorf("error when finding user by email: %w", err)
	}
	return user, nil
}
