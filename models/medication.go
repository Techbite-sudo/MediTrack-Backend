// medication.go
package models

import (
	"time"
)

type Medication struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
