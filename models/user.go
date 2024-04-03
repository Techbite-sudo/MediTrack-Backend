package models

import (
	"time"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
