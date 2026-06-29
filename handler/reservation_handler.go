package handler

import (
	"net/http"
	"strconv"

	"github.com/Mokarama/assignment-6-spotsync-api/dto"
	"github.com/Mokarama/assignment-6-spotsync-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	service  *service.ReservationService
	validate *validator.Validate
}

func NewReservationHandler() *ReservationHandler {
	return &ReservationHandler{
		service:  service.NewReservationService(),
		validate: validator.New(),
	}
}

// Create Reservation
func (h *ReservationHandler) Create(c echo.Context) error {

	var req dto.CreateReservationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	// Get User ID from JWT
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userID := uint(claims["user_id"].(float64))

	if err := h.service.Create(userID, &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Reservation created successfully",
	})
}

// Get All Reservations
func (h *ReservationHandler) GetAll(c echo.Context) error {

	reservations, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, reservations)
}

// Get Reservation By ID
func (h *ReservationHandler) GetByID(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID",
		})
	}

	reservation, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Reservation not found",
		})
	}

	return c.JSON(http.StatusOK, reservation)
}

// Cancel Reservation
func (h *ReservationHandler) Cancel(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID",
		})
	}

	if err := h.service.Cancel(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Reservation cancelled successfully",
	})
}

// Delete Reservation
func (h *ReservationHandler) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID",
		})
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Reservation deleted successfully",
	})
}
