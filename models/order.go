package models

import (
    "time"
)

type Order struct {
    ID        string     `gorm:"primaryKey" json:"id"`
    UserID    string     `gorm:"size:36" json:"userId"` // Change to varchar(36)
    User      User       `gorm:"foreignKey:UserID" json:"user"`
    Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
    Total     float64    `json:"total"`
    Status    string     `json:"status"`
    CreatedAt time.Time  `json:"createdAt"`
    UpdatedAt time.Time  `json:"updatedAt"`
}

type OrderItem struct {
    ID           string      `gorm:"primaryKey" json:"id"`
    OrderID      string      `gorm:"size:36" json:"orderId"`
    MedicationID string      `gorm:"size:36" json:"medicationId"`
    Medication   Medication  `gorm:"foreignKey:MedicationID" json:"medication"`
    Quantity     int         `json:"quantity"`
    Price        float64     `json:"price"`
}