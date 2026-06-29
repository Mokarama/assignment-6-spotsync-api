package routes

import (
	"github.com/Mokarama/assignment-6-spotsync-api/handler"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo) {
	authHandler := handler.NewAuthHandler()

	auth := e.Group("/api/v1/auth")

	// Authentication Routes
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
}
