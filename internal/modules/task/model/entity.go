package task_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Status      string    `gorm:"type:varchar(255);not null"`
	UsedID      uuid.UUID `gorm:"not null"`
}
