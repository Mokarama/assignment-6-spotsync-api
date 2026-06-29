package repository

import (
	"github.com/Mokarama/assignment-6-spotsync-api/database"
	"github.com/Mokarama/assignment-6-spotsync-api/models"
)

type ParkingZoneRepository struct{}

func NewParkingZoneRepository() *ParkingZoneRepository {
	return &ParkingZoneRepository{}
}

// Create Parking Zone
func (r *ParkingZoneRepository) Create(zone *models.ParkingZone) error {
	return database.DB.Create(zone).Error
}

// Get All Parking Zones
func (r *ParkingZoneRepository) GetAll() ([]models.ParkingZone, error) {
	var zones []models.ParkingZone

	err := database.DB.Find(&zones).Error
	if err != nil {
		return nil, err
	}

	// Calculate Available Spots
	for i := range zones {

		var activeReservations int64

		err := database.DB.Model(&models.Reservation{}).
			Where("zone_id = ? AND status = ?", zones[i].ID, "active").
			Count(&activeReservations).Error

		if err != nil {
			return nil, err
		}

		zones[i].AvailableSpots = zones[i].TotalCapacity - int(activeReservations)
	}

	return zones, nil
}

// Get Parking Zone By ID
func (r *ParkingZoneRepository) GetByID(id uint) (*models.ParkingZone, error) {
	var zone models.ParkingZone

	err := database.DB.First(&zone, id).Error
	if err != nil {
		return nil, err
	}

	var activeReservations int64

	err = database.DB.Model(&models.Reservation{}).
		Where("zone_id = ? AND status = ?", zone.ID, "active").
		Count(&activeReservations).Error

	if err != nil {
		return nil, err
	}

	zone.AvailableSpots = zone.TotalCapacity - int(activeReservations)

	return &zone, nil
}

// Update Parking Zone
func (r *ParkingZoneRepository) Update(zone *models.ParkingZone) error {
	return database.DB.Save(zone).Error
}

// Delete Parking Zone
func (r *ParkingZoneRepository) Delete(id uint) error {
	return database.DB.Delete(&models.ParkingZone{}, id).Error
}
