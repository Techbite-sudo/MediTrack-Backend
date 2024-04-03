package users

import (
	"github.com/Techbite-sudo/MediTrack-Backend/database"
	"github.com/Techbite-sudo/MediTrack-Backend/graph/model"
	"github.com/google/uuid"
	"time"
)

func CreateUser(input model.CreateUserInput) (*model.User, error) {
	
	id := uuid.New()

	user := model.User{
		ID:        id.String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, input *model.UpdateUserInput) (*model.User, error) {
	var user model.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Phone != nil && *input.Phone != "" {
		user.Phone = input.Phone
	}
	if input.Address != nil && *input.Address != "" {
		user.Address = input.Address
	}

	err = database.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FetchUsers() ([]*model.User, error) {
	var users []*model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err

}

func FetchUser(id string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
