package product

import (
	"errors"
	"learn_cruds_go/config"
	"learn_cruds_go/models"
	"net/http"

	"gorm.io/gorm"
)

func fetchAllProducts(userID uint) (int, []models.Product) {
	var products []models.Product

	result := config.DB.Model(&products).Where("user_id = ?", userID).Find(&products)

	if result.Error != nil {
		return http.StatusInternalServerError, products
	}

	return http.StatusOK, products
}

func fetchProductById(productId uint) (int, models.Product) {
    var product models.Product

    // Mencari berdasarkan ID (Primary Key) secara langsung
    result := config.DB.First(&product, productId)

    if result.Error != nil {
        // Cek apakah errornya karena data memang tidak ada
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return http.StatusNotFound, product
        }
        // Jika error selain itu (misal: DB mati), baru kirim 500
        return http.StatusInternalServerError, product
    }

    return http.StatusOK, product
}

func addProduct(newProduct models.Product) (int, string) {
	result := config.DB.Create(&newProduct)

	if result.Error != nil {
		return http.StatusInternalServerError, "Failed to insert product"
	}

	return http.StatusCreated, "Successfully create product"
}

func updateProduct(newProductData models.Product) (int, models.Product) {
	var product models.Product
	
	result := config.DB.Model(&product).Where("ID = ?", newProductData.ID).Select("ProductName", "Price").Updates(newProductData)

	if result.Error != nil {
		return http.StatusInternalServerError, product
	}

	return http.StatusOK, product
}

func deleteProduct(productId uint) (int) {
	result := config.DB.Unscoped().Delete(&models.Product{}, productId)

	if result.Error != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}