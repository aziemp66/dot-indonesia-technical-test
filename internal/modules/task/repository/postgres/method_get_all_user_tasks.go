package task_repository_postgres

import (
	"context"
	"fmt"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	"github.com/google/uuid"
)

func (r *taskRepositoryPostgres) GetAllUserTasks(ctx context.Context, userID uuid.UUID) ([]task_model.Task, error) {
	var tasks []task_model.Task

	if err := r.db.Where(&task_model.Task{UsedID: userID}).Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("error when finding user by email: %w", err)
	}

	return tasks, nil
}
