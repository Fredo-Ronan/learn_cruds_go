package user

import (
	"fmt"
	"learn_cruds_go/config"
	"learn_cruds_go/models"
	"net/http"
)

func fetchUserData(userID uint) (int, models.User) {
	var foundUser models.User

	result := config.DB.Where("ID = ?", userID).First(&foundUser)

	if result.Error != nil {
		return http.StatusNotFound, foundUser
	}

	return http.StatusOK, foundUser
}

func updateUserData(newUserData *models.User, userID uint) (int, models.User) {
	var existingEmail models.User

	// check for duplicate email
	config.DB.Where("email = ? AND id != ?", newUserData.Email, userID).First(&existingEmail)

	if existingEmail.ID != 0 {
		return http.StatusConflict, existingEmail
	}

	var user models.User

	result := config.DB.Model(&user).Where("ID = ?", userID).Select("Username", "Email").Updates(newUserData)

	fmt.Println("UPDATE OK!")

	if result.Error != nil {
		return http.StatusInternalServerError, user
	}

	return fetchUserData(userID)
}
