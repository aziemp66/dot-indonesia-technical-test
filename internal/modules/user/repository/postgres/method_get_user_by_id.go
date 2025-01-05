package user_repository_postgres

import (
	"context"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"

	"github.com/google/uuid"
)

func (r *userRepositoryPostgres) GetUserByID(ctx context.Context, id uuid.UUID) (user_model.User, error) {
	user := user_model.User{
		ID: id,
	}
	if err := r.db.First(&user).Error; err != nil {
		return user, fmt.Errorf("error when finding user by ID: %w", err)
	}
	return user, nil
}
