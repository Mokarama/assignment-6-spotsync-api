package routes

import (
	"github.com/Mokarama/assignment-6-spotsync-api/handler"
	"github.com/Mokarama/assignment-6-spotsync-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterReservationRoutes(e *echo.Echo) {

	reservationHandler := handler.NewReservationHandler()

	reservations := e.Group("/api/v1/reservations")

	// Protected Routes
	reservations.Use(middleware.JWTMiddleware)

	reservations.POST("", reservationHandler.Create)
	reservations.GET("", reservationHandler.GetAll)
	reservations.GET("/:id", reservationHandler.GetByID)
	reservations.PATCH("/:id/cancel", reservationHandler.Cancel)
	reservations.DELETE("/:id", reservationHandler.Delete)
}
