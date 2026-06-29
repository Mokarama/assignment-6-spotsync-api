package repository

import (
	"github.com/Mokarama/assignment-6-spotsync-api/database"
	"github.com/Mokarama/assignment-6-spotsync-api/models"
)

type ReservationRepository struct{}

func NewReservationRepository() *ReservationRepository {
	return &ReservationRepository{}
}

// Create Reservation
func (r *ReservationRepository) Create(reservation *models.Reservation) error {
	return database.DB.Create(reservation).Error
}

// Get All Reservations
func (r *ReservationRepository) GetAll() ([]models.Reservation, error) {
	var reservations []models.Reservation

	err := database.DB.Preload("User").Preload("Zone").Find(&reservations).Error

	return reservations, err
}

// Get Reservation By ID
func (r *ReservationRepository) GetByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation

	err := database.DB.Preload("User").Preload("Zone").First(&reservation, id).Error

	return &reservation, err
}

// Update Reservation
func (r *ReservationRepository) Update(reservation *models.Reservation) error {
	return database.DB.Save(reservation).Error
}

// Delete Reservation
func (r *ReservationRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Reservation{}, id).Error
}
