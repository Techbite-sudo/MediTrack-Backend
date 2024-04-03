package models

import (
	"time"
	uuid "github.com/satori/go.uuid"
)

type Medication struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
