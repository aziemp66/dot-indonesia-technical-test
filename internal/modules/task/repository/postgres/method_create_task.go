package task_repository_postgres

import (
	"context"
	"fmt"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	"github.com/google/uuid"
)

func (r *taskRepositoryPostgres) CreateTask(ctx context.Context, userID uuid.UUID, title string, description string, status string) (id string, err error) {
	task := task_model.Task{
		Title:       title,
		Description: description,
		Status:      status,
		UsedID:      userID,
	}

	if err := r.db.WithContext(ctx).Create(&task).Error; err != nil {
		return "", fmt.Errorf("error when creating user: %w", err)
	}
	return task.ID.String(), nil
}
