package migrations

import (
	"log"
	"github.com/Techbite-sudo/MediTrack-Backend/database"
	"github.com/Techbite-sudo/MediTrack-Backend/models"
)

func Migrate() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Medication{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}