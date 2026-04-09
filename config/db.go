package config

import (
	"learn_cruds_go/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database!")
	}

	database.AutoMigrate(&models.User{}, &models.Product{})

	DB = database
}