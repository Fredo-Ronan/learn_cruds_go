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

func GetUserDataHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	status, result := fetchUserData(userID)

	c.HTML(status, "user_data_component", gin.H{
		"Username": result.Username,
		"Email":    result.Email,
	})
}

func GetUserEditHandler(c *gin.Context) {
	userID := c.MustGet(config.SessionCurrentUser).(uint)

	status, result := fetchUserData(userID)

	c.HTML(status, "user_data_edit_component", gin.H{
		"Username": result.Username,
		"Email":    result.Email,
	})
}

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
