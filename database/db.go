package database

import (
    "fmt"
    "github.com/Techbite-sudo/MediTrack-Backend/config"
    "github.com/Techbite-sudo/MediTrack-Backend/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) error {
    connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.DBUser,
        config.DBPassword,
        config.DBHost,
        config.DBPort,
        config.DBName,
    )

    db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
    if err != nil {
        return err
    }

    DB = db

    // Migrate the schema
    err = DB.AutoMigrate(&models.User{}, &models.Medication{}, &models.Order{}, &models.OrderItem{})
    if err != nil {
        log.Println("Error migrating schema:", err)
        return err
    }

    return nil
}