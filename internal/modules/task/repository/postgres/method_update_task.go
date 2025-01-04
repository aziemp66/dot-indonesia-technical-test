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

func (r *taskRepositoryPostgres) UpdateTask(ctx context.Context, id uuid.UUID, title string, description string, status string) error {
	task := task_model.Task{
		Title:       title,
		Description: description,
		Status:      status,
	}

	if err := r.db.WithContext(ctx).Model(&task_model.Task{}).Where("id = ?", id).Updates(task).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return util_error.NewNotFound(err, "task not found")
	} else if err != nil {
		return fmt.Errorf("error when updating task: %w", err)
	}

	return nil
}
