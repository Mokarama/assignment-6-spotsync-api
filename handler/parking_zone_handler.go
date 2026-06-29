package handler

import (
	"net/http"
	"strconv"

	"github.com/Mokarama/assignment-6-spotsync-api/dto"
	"github.com/Mokarama/assignment-6-spotsync-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ParkingZoneHandler struct {
	service  *service.ParkingZoneService
	validate *validator.Validate
}

func NewParkingZoneHandler() *ParkingZoneHandler {
	return &ParkingZoneHandler{
		service:  service.NewParkingZoneService(),
		validate: validator.New(),
	}
}

// Create Parking Zone
func (h *ParkingZoneHandler) Create(c echo.Context) error {
	var req dto.CreateParkingZoneRequest

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

	if err := h.service.Create(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Parking zone created successfully",
	})
}

// Get All Parking Zones
func (h *ParkingZoneHandler) GetAll(c echo.Context) error {
	zones, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, zones)
}

// Get Parking Zone By ID
func (h *ParkingZoneHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID",
		})
	}

	zone, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Parking zone not found",
		})
	}

	return c.JSON(http.StatusOK, zone)
}

// Update Parking Zone
func (h *ParkingZoneHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID",
		})
	}

	var req dto.UpdateParkingZoneRequest

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

	if err := h.service.Update(uint(id), &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Parking zone updated successfully",
	})
}

// Delete Parking Zone
func (h *ParkingZoneHandler) Delete(c echo.Context) error {
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
		"message": "Parking zone deleted successfully",
	})
}
