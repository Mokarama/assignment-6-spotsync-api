package service

import (
	"github.com/Mokarama/assignment-6-spotsync-api/dto"
	"github.com/Mokarama/assignment-6-spotsync-api/models"
	"github.com/Mokarama/assignment-6-spotsync-api/repository"
)

type ParkingZoneService struct {
	repo *repository.ParkingZoneRepository
}

func NewParkingZoneService() *ParkingZoneService {
	return &ParkingZoneService{
		repo: repository.NewParkingZoneRepository(),
	}
}

// Create Parking Zone
func (s *ParkingZoneService) Create(req *dto.CreateParkingZoneRequest) error {

	zone := models.ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	return s.repo.Create(&zone)
}

// Get All Parking Zones
func (s *ParkingZoneService) GetAll() ([]models.ParkingZone, error) {
	return s.repo.GetAll()
}

// Get Parking Zone By ID
func (s *ParkingZoneService) GetByID(id uint) (*models.ParkingZone, error) {
	return s.repo.GetByID(id)
}

// Update Parking Zone
func (s *ParkingZoneService) Update(id uint, req *dto.UpdateParkingZoneRequest) error {

	zone, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	zone.Name = req.Name
	zone.Type = req.Type
	zone.TotalCapacity = req.TotalCapacity
	zone.PricePerHour = req.PricePerHour

	return s.repo.Update(zone)
}

// Delete Parking Zone
func (s *ParkingZoneService) Delete(id uint) error {
	return s.repo.Delete(id)
}
