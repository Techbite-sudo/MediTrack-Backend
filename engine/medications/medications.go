package medications

import (
	"github.com/Techbite-sudo/MediTrack-Backend/database"
	"github.com/Techbite-sudo/MediTrack-Backend/graph/model"
)

func FetchMedications() ([]*model.Medication, error) {
	var medications []*model.Medication
	err := database.DB.Find(&medications).Error
	if err != nil {
		return nil, err
	}
	return medications, nil
}

func FetchMedication(id string) (*model.Medication, error) {
	var medication model.Medication
	err := database.DB.Where("id = ?", id).First(&medication).Error
	if err != nil {
		return nil, err
	}
	return &medication, nil
}
