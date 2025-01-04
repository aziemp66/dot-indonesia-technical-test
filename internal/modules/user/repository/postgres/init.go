package user_repository_postgres

import (
	user_repository "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/repository"

	"gorm.io/gorm"
)

type userRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepositoryPostgres(db *gorm.DB) user_repository.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}
