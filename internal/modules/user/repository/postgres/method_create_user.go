package user_repository_postgres

import (
	"context"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"

	"github.com/google/uuid"
)

func (r *userRepositoryPostgres) CreateUser(ctx context.Context, email, hashedPassword, name string) (string, error) {
	user := &user_model.User{
		ID:       uuid.New(),
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	}
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return "", fmt.Errorf("error when creating user: %w", err)
	}
	return user.ID.String(), nil
}
