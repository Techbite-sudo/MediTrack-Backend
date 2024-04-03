package models

import (
    "time"
)

type User struct {
    ID        string    `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name"`
    Email     string    `gorm:"unique" json:"email"`
    Password  string    `json:"-"`
    Phone     string    `json:"phone"`
    Address   string    `json:"address"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}