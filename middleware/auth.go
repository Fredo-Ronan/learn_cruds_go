package middleware

import (
	"learn_cruds_go/config"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := config.Store.Get(c.Request, config.SessionName)

		if uid, ok := session.Values[config.SessionUserID].(uint); ok {
			c.Set(config.SessionCurrentUser, uid)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Login Required"})
		c.Abort()
	}
}