package task_repository_postgres

import (
	task_repository "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/repository"
	"gorm.io/gorm"
)

type taskRepositoryPostgres struct {
	db *gorm.DB
}

func NewTaskRepositoryPostgres(db *gorm.DB) task_repository.TaskRepository {
	return &taskRepositoryPostgres{
		db: db,
	}
}
