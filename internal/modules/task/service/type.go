package task_service

import (
	"context"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
)

type TaskService interface {
	CreateTask(ctx context.Context, userID string, title, description, status string) (id string, err error)
	GetTaskByID(ctx context.Context, id string) (task_model.GetTaskResponse, error)
	GetAllUserTasks(ctx context.Context, userID string) ([]task_model.GetTaskResponse, error)
	UpdateTask(ctx context.Context, id string, title, description, status string) error
	DeleteTask(ctx context.Context, id string) error
}
