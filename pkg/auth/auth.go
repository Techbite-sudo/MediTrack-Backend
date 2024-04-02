package auth

import (
	"errors"
	"github.com/Techbite-sudo/MediTrack-Backend/database"
	"github.com/Techbite-sudo/MediTrack-Backend/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Authenticate(email, password string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
