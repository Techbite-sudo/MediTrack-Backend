package database

import (
	"fmt"
	"github.com/Techbite-sudo/MediTrack-Backend/config"
	"github.com/Techbite-sudo/MediTrack-Backend/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) error {
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(sqlserver.Open(connStr), &gorm.Config{})
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
