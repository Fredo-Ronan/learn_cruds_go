package user

import (
	"learn_cruds_go/config"
	"learn_cruds_go/middleware"
	"learn_cruds_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/user")
	userGroup.Use(middleware.Auth())
	{
		userGroup.GET("/me", GetUserDataHandler)
		userGroup.GET("/edit", GetUserEditHandler)
		userGroup.PUT("/update", UpdateUserDataHandler)
	}
}

// GetUserDataHandler godoc
// @Summary			Get user data
// @Description		Route ini digunakan ketika client ingin mengambil data user yang sedang login sesuai data di session
// @Tags			user
// @Security		CookieAuth
// @Produce			html
// @Success      	200
// @Router			/v1/user/me [get]
func GetUserDataHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	status, result := fetchUserData(userID)

	c.HTML(status, "user_data_component", gin.H{
		"Username": result.Username,
		"Email":    result.Email,
	})
}

// GetUserEditHandler godoc
// @Summary			Get user edit component
// @Description		Route ini digunakan ketika client ingin mengedit data sehingga backend mengembalikan tampilan edit
// @Tags			user
// @Security		CookieAuth
// @Produce			html
// @Success      	200
// @Router			/v1/user/edit [get]
func GetUserEditHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	status, result := fetchUserData(userID)

	c.HTML(status, "user_data_edit_component", gin.H{
		"Username": result.Username,
		"Email":    result.Email,
	})
}

// UpdateUserDataHandler godoc
// @Summary			Update user data
// @Description		Route ini digunakan ketika client ingin mengedit dan menyimpan data user terbaru
// @Tags			user
// @Security		CookieAuth
// @Produce			html
// @Param			username		formData		string		true 		"Username user"
// @Param			email			formData		string		true		"Email user"
// @Success      	200
// @Router			/v1/user/update [put]
func UpdateUserDataHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	var newUserData models.User

	if err := c.ShouldBind(&newUserData); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	status, result := updateUserData(&newUserData, userID)

	c.HTML(status, "user_data_component", gin.H{
		"Username": result.Username,
		"Email":    result.Email,
	})
}
