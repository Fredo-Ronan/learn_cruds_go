package auth

import (
	"fmt"
	"learn_cruds_go/config"
	"learn_cruds_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", RegisterUserHandler)
		authGroup.POST("/login", LoginHandler)
		authGroup.POST("/logout", LogoutHandler)
	}
}

// RegisterUserHandler godoc
// @Summary			Register atau buat data user baru
// @Description		Route ini akan membuat data user baru
// @Tags			auth
// @Accept       	x-www-form-urlencoded
// @Param        	email     formData  string  true  "Email User"
// @Param        	password  formData  string  true  "Password User"
// @Success      	200       {string} string "Successfully registered user"
// @Failure			500		  {string} string "Internal Server Error"
// @Router			/v1/auth/register [post]
func RegisterUserHandler(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBind(&newUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	status, result := RegisterUser(&newUser)

	c.String(status, result)
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginHandler godoc
// @Summary			User Login
// @Description		Route ini dipanggil ketika ada user ingin login
// @Tags			auth
// @Accept       	x-www-form-urlencoded
// @Param        	email     formData  string  true  "Email User"
// @Param        	password  formData  string  true  "Password User"
// @Success      	302       {string} string "Redirected to /dashboard"
// @Failure			401		  {string} string "Email atau Password Salah!"
// @Router			/v1/auth/login [post]
func LoginHandler(c *gin.Context) {
	var loginData Login

	if err := c.ShouldBind(&loginData); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	status, result := LoginUser(loginData.Email, loginData.Password)

	if status == http.StatusUnauthorized {
		c.String(status, "Email atau Password Salah!")
		return
	}

	fmt.Println("LOGIN OK!")

	// create session here
	session, _ := config.Store.Get(c.Request, config.SessionName)
	session.Values[config.SessionUserID] = result.ID
	session.Values[config.SessionAuthenticated] = true
	session.Save(c.Request, c.Writer)

	// redirect to dashboard
	// if using HTMX
	c.Header("HX-Redirect", "/dashboard")

	// if using HTML
	// c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}

// LogoutHandler godoc
// @Summary			User Logout
// @Description		Route ini dipanggil ketika user ingin logout. Route ini akan mengambil data dari session currentUser dan membuat maxAge -1 untuk otomatis menghapusnya dari browser (invalidate)
// @Tags			auth
// @Success      	302       {string} string "Redirected to /login"
// @Failure			500		  {string} string "Internal Server Error"
// @Router			/v1/auth/logout [post]
func LogoutHandler(c *gin.Context) {
	session, _ := config.Store.Get(c.Request, config.SessionName)

	session.Options.MaxAge = -1

	session.Values["authenticated"] = false
	session.Values["userID"] = nil

	err := session.Save(c.Request, c.Writer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to logout")
		return
	}

	// if using HTMX
	c.Header("HX-Redirect", "/login")

	// if using HTML
	// c.Redirect(http.StatusTemporaryRedirect, "/login")
}
