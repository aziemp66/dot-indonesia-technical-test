package task_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *taskRepositoryPostgres) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Delete(&task_model.Task{
		ID: id,
	}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return util_error.NewNotFound(err, "task not found")
	} else if err != nil {
		return fmt.Errorf("error when deleting task: %w", err)
	}

	return nil
}
