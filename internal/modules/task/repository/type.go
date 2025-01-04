package task_repository

import (
	"context"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	"github.com/google/uuid"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, userID uuid.UUID, title, description, status string) (id string, err error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (task_model.Task, error)
	GetAllUserTasks(ctx context.Context, userID uuid.UUID) ([]task_model.Task, error)
	UpdateTask(ctx context.Context, id uuid.UUID, title, description, status string) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
}
