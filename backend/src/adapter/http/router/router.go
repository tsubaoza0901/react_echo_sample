package router

import (
	"os"
	"react-echo-sample/adapter/http/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var signingKey = []byte(os.Getenv("SIGNINGKEY"))

// Config ...
var Config = middleware.JWTConfig{
	SigningKey: signingKey,
}

// InitRouting ...
func InitRouting(e *echo.Echo, h handler.AppHandler) {
	apiBaseURL := "/api/v1/"

	// Auth
	e.POST(apiBaseURL+"signup", h.Signup)
	e.POST(apiBaseURL+"login", h.Login)

	// 認証が必要なルーティングのグループ化
	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(Config))

	// User
	// e.GET(apiBaseURL+"user/:id", h.GetUser)
	// e.GET(apiBaseURL+"users", h.GetUsers)
	auth.PUT(apiBaseURL+"user/:id", h.UpdateUser)
	// e.DELETE(apiBaseURL+"user/:id", h.DeleteUser)
}
