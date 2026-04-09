package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username" form:"username"`
	Password string    `json:"password" form:"password" binding:"required" gorm:"not null"`
	Email    string    `json:"email" form:"email" binding:"required" gorm:"unique; not null"`
	Products []Product `json:"products" gorm:"foreignKey:UserID"` // one-to-many relation with Product
}