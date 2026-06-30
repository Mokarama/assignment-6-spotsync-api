package repository

import (
	"errors"

	"github.com/Mokarama/assignment-6-spotsync-api/database"
	"github.com/Mokarama/assignment-6-spotsync-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReservationRepository struct{}

func NewReservationRepository() *ReservationRepository {
	return &ReservationRepository{}
}

// Create Reservation (Transaction + Row Lock)
func (r *ReservationRepository) CreateWithTransaction(reservation *models.Reservation) error {

	return database.DB.Transaction(func(tx *gorm.DB) error {

		// Lock Parking Zone Row
		var zone models.ParkingZone

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&zone, reservation.ZoneID).Error; err != nil {
			return err
		}

		// Count Active Reservations
		var activeReservations int64

		if err := tx.Model(&models.Reservation{}).
			Where("zone_id = ? AND status = ?", reservation.ZoneID, "active").
			Count(&activeReservations).Error; err != nil {
			return err
		}

		// Check Capacity
		if activeReservations >= int64(zone.TotalCapacity) {
			return errors.New("parking zone is full")
		}

		// Create Reservation
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}

		return nil
	})
}

// Get All Reservations
func (r *ReservationRepository) GetAll() ([]models.Reservation, error) {
	var reservations []models.Reservation

	err := database.DB.
		Preload("User").
		Preload("Zone").
		Find(&reservations).Error

	return reservations, err
}

// Get Reservation By ID
func (r *ReservationRepository) GetByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation

	err := database.DB.
		Preload("User").
		Preload("Zone").
		First(&reservation, id).Error

	return &reservation, err
}

// Get Reservations By User ID
func (r *ReservationRepository) GetByUserID(userID uint) ([]models.Reservation, error) {
	var reservations []models.Reservation

	err := database.DB.
		Preload("User").
		Preload("Zone").
		Where("user_id = ?", userID).
		Find(&reservations).Error

	return reservations, err
}

// Update Reservation
func (r *ReservationRepository) Update(reservation *models.Reservation) error {
	return database.DB.Save(reservation).Error
}

// Delete Reservation
func (r *ReservationRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Reservation{}, id).Error
}
