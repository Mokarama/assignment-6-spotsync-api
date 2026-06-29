package routes

import (
	"github.com/Mokarama/assignment-6-spotsync-api/handler"
	"github.com/Mokarama/assignment-6-spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterParkingZoneRoutes(e *echo.Echo) {

	parkingZoneHandler := handler.NewParkingZoneHandler()

	// ==========================
	// Public Routes
	// ==========================
	e.GET("/api/v1/zones", parkingZoneHandler.GetAll)
	e.GET("/api/v1/zones/:id", parkingZoneHandler.GetByID)

	// ==========================
	// Admin Protected Routes
	// ==========================
	zones := e.Group("/api/v1/zones")

	// JWT Authentication
	zones.Use(middleware.JWTMiddleware)

	// Admin Authorization
	zones.Use(middleware.AdminMiddleware)

	zones.POST("", parkingZoneHandler.Create)
	zones.PATCH("/:id", parkingZoneHandler.Update)
	zones.DELETE("/:id", parkingZoneHandler.Delete)
}
