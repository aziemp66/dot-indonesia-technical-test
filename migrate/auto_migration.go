package migrate

import (
	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&user_model.User{}, &task_model.Task{})
}
