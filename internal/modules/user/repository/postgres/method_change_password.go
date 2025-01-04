package user_repository_postgres

import (
	"context"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"

	"github.com/google/uuid"
)

func (r *userRepositoryPostgres) ChangePassword(ctx context.Context, id uuid.UUID, hashedPassword string) error {
	if err := r.db.WithContext(ctx).Model(&user_model.User{}).Where("id = ?", id).Update("password", hashedPassword).Error; err != nil {
		return fmt.Errorf("error when updating password: %w", err)
	}
	return nil
}
