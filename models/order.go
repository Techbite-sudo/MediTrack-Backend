package models

import (
	"time"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User        `gorm:"foreignKey:UserID" json:"user"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	OrderID      uuid.UUID `gorm:"type:uuid"`
	MedicationID uuid.UUID `gorm:"type:uuid"`
	Medication   Medication `gorm:"foreignKey:MedicationID" json:"medication"`
	Quantity     int        `json:"quantity"`
	Price        float64    `json:"price"`
}
