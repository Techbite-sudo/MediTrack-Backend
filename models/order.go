package models

import (
	"time"
)

type Order struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	UserID    uint        `json:"userId"`
	User      User        `gorm:"foreignKey:UserID" json:"user"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type OrderItem struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	OrderID      uint       `json:"orderId"`
	MedicationID uint       `json:"medicationId"`
	Medication   Medication `gorm:"foreignKey:MedicationID" json:"medication"`
	Quantity     int        `json:"quantity"`
	Price        float64    `json:"price"`
}
