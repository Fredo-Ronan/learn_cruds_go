package main

import (
	"learn_cruds_go/config"
	"learn_cruds_go/middleware"
	"learn_cruds_go/v1/auth"
	"learn_cruds_go/v1/product"
	"learn_cruds_go/v1/user"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load() // load .env file

	if errEnv != nil {
		log.Fatal(errEnv)
	}

	r := gin.Default()

	// 1. Kumpulkan semua file .html secara recursive
	var templates []string
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Ambil hanya file yang ekstensinya .html
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	})

	if err != nil {
		panic("Gagal men-scan folder templates: " + err.Error())
	}

	// 2. Load semua file yang terkumpul ke Gin
	r.LoadHTMLFiles(templates...)

	// unprotected routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// protected routes
	// install middleware here
	r.GET("/dashboard", middleware.Auth(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	config.ConnectDatabase() // initialize database SQLite
	config.InitSession()     // initialize Store object for session management

	v1 := r.Group("/v1")
	{
		auth.RegisterAuthRoutes(v1)
		user.RegisterUserRoutes(v1)
		product.RegisterProductRoutes(v1)
	}

	r.Run(":8000")
}
