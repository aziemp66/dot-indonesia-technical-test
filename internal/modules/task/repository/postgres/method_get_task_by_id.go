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

func (r *taskRepositoryPostgres) GetTaskByID(ctx context.Context, id uuid.UUID) (task_model.Task, error) {
	task := task_model.Task{
		ID: id,
	}

	if err := r.db.First(&task).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return task, util_error.NewNotFound(err, "task not found")
	} else if err != nil {
		return task, fmt.Errorf("error when finding task by ID: %w", err)
	}

	return task, nil
}
