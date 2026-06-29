package routes

import (
	"github.com/Mokarama/assignment-6-spotsync-api/handler"
	"github.com/Mokarama/assignment-6-spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterReservationRoutes(e *echo.Echo) {

	reservationHandler := handler.NewReservationHandler()

	// ==================================
	// Authenticated Driver Routes
	// ==================================
	driver := e.Group("/api/v1/reservations")
	driver.Use(middleware.JWTMiddleware)

	driver.POST("", reservationHandler.Create)
	driver.GET("/:id", reservationHandler.GetByID)
	driver.PATCH("/:id/cancel", reservationHandler.Cancel)

	// (এই Route আমরা পরের ধাপে implement করব)
	driver.GET("/my-reservations", reservationHandler.GetMyReservations)

	// ==================================
	// Admin Routes
	// ==================================
	admin := e.Group("/api/v1/reservations")
	admin.Use(middleware.JWTMiddleware)
	admin.Use(middleware.AdminMiddleware)

	admin.GET("", reservationHandler.GetAll)
	admin.DELETE("/:id", reservationHandler.Delete)
}
