package user_repository_postgres

import (
	"context"
	"fmt"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"

	"github.com/google/uuid"
)

func (r *userRepositoryPostgres) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Delete(&user_model.User{}, id).Error; err != nil {
		return fmt.Errorf("error when deleting user: %w", err)
	}
	return nil
}
