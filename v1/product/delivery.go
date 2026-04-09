package product

import (
	"learn_cruds_go/config"
	"learn_cruds_go/middleware"
	"learn_cruds_go/models"
	"learn_cruds_go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup) {
	productGroup := rg.Group("/product")
	productGroup.Use(middleware.Auth())
	{
		productGroup.GET("/all", GetAllProductHandler)
		productGroup.GET("/get/:id", GetProductByIdHandler)
		productGroup.GET("/add/component", GetAddProductHandler)
		productGroup.POST("/add", AddProductHandler)

		productGroup.GET("/edit/component/:id", GetEditProductHandler)
		productGroup.PUT("/edit", EditProductHandler)

		productGroup.DELETE("/delete/:id", DeleteProductHandler)
	}
}

func GetAllProductHandler(c *gin.Context){
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	status, result := fetchAllProducts(userID)

	c.HTML(status, "product_list", gin.H{
		"Products": result,
	})
}

func GetProductByIdHandler(c *gin.Context){
	productId := c.Param("id")

	// parse to unit64
	u64ProductID, err := strconv.ParseUint(productId, 10, 32)

	if err != nil {
		c.String(http.StatusInternalServerError, "Parsing failed! Internal Server Error!")
		return
	}

	// parse to uint
	uProductID := uint(u64ProductID)

	status, result := fetchProductById(uProductID)

	c.HTML(status, "product_row_view", result)
}

func GetAddProductHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	status, result := fetchAllProducts(userID)

	c.HTML(status, "product_add_component", gin.H{
		"Products": result,
	})
}

func GetEditProductHandler(c *gin.Context) {
	productId := c.Param("id")

	// parse to uint
	uProductID, err := utils.ConvertUint(productId)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	status, result := fetchProductById(uProductID)

	c.HTML(status, "product_edit_component", result)
}

func AddProductHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	var newProduct models.Product

	if err := c.ShouldBind(&newProduct); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	newFinalProduct := models.Product{
		ProductName: newProduct.ProductName,
		Price: newProduct.Price,
		UserID: userID,
	}

	status, _ := addProduct(newFinalProduct)
	_, products := fetchAllProducts(userID)

	c.HTML(status, "product_list", gin.H{
		"Products": products,
	})
}

func EditProductHandler(c *gin.Context){
	var newProductData models.Product

	if err := c.ShouldBind(&newProductData); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	status, result := updateProduct(newProductData)

	c.HTML(status, "product_row_view", result)
}

func DeleteProductHandler(c *gin.Context){
	productId := c.Param("id")

	uProductID, err := utils.ConvertUint(productId)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	status := deleteProduct(uProductID)

	c.Status(status)
}