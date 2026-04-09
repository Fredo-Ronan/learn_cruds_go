package models

import "gorm.io/gorm"
import "golang.org/x/text/language"
import "golang.org/x/text/message"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name" form:"product_name" binding:"required"`
	Price int `json:"price" form:"price" binding:"required"`
	UserID uint `json:"user_id"` // foreign key
}

func (p Product) FormatPrice() string {
    printer := message.NewPrinter(language.Indonesian)
    return printer.Sprintf("Rp. %d,00", p.Price)
}