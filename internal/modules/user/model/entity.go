package user_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Email    string    `gorm:"type:varchar(255);not null;unique"`
	Password string    `gorm:"type:varchar(255);not null"`
	Name     string    `gorm:"type:varchar(255);not null"`
}
