package auth

import (
	"learn_cruds_go/config"
	"learn_cruds_go/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(newUser *models.User) (int, string) {
	// hash the password first
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)

	securedUser := models.User{
		Email:    newUser.Email,
		Password: string(hashedPassword),
		Username: newUser.Username,
	}

	result := config.DB.Create(&securedUser)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error.Error()
	}

	return http.StatusCreated, "Successfully registered user"
}

func LoginUser(email string, password string) (int, models.User) {
	var foundUser models.User

	// find the user
	result := config.DB.Where("email = ?", email).First(&foundUser)

	if result.Error != nil {
		return http.StatusNotFound, foundUser
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))

	if errPassword != nil {
		return http.StatusUnauthorized, foundUser
	}

	return http.StatusOK, foundUser
}
