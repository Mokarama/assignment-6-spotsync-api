package routes

import (
	"github.com/Mokarama/assignment-6-spotsync-api/handler"
	"github.com/Mokarama/assignment-6-spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterParkingZoneRoutes(e *echo.Echo) {
	parkingZoneHandler := handler.NewParkingZoneHandler()

	zones := e.Group("/api/v1/zones")

	// Protected Routes
	zones.Use(middleware.JWTMiddleware)

	zones.POST("", parkingZoneHandler.Create)
	zones.GET("", parkingZoneHandler.GetAll)
	zones.GET("/:id", parkingZoneHandler.GetByID)
	zones.PATCH("/:id", parkingZoneHandler.Update)
	zones.DELETE("/:id", parkingZoneHandler.Delete)
}
