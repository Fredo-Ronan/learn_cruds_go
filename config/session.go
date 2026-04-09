package config

import (
	"os"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "test-app-session"
	SessionUserID = "userID"
	SessionAuthenticated = "authenticated"
	SessionCurrentUser = "current_user"
)

var Store *sessions.CookieStore

func InitSession() {
	key := os.Getenv("SECRET_TOKEN")

	Store = sessions.NewCookieStore([]byte(key))

	Store.Options = &sessions.Options{
		Path: "/",
		MaxAge: 3600 * 24, // 1 day
		HttpOnly: true,
	}
}