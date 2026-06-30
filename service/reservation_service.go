package service

import (
	"errors"

	"github.com/Mokarama/assignment-6-spotsync-api/dto"
	"github.com/Mokarama/assignment-6-spotsync-api/models"
	"github.com/Mokarama/assignment-6-spotsync-api/repository"
)

type ReservationService struct {
	repo *repository.ReservationRepository
}

func NewReservationService() *ReservationService {
	return &ReservationService{
		repo: repository.NewReservationRepository(),
	}
}

// Create Reservation
func (s *ReservationService) Create(userID uint, req *dto.CreateReservationRequest) error {

	reservation := models.Reservation{
		UserID:       userID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
		Status:       "active",
	}

	// Transaction + Row Lock
	return s.repo.CreateWithTransaction(&reservation)
}

// Get All Reservations
func (s *ReservationService) GetAll() ([]models.Reservation, error) {
	return s.repo.GetAll()
}

// Get Reservation By ID
func (s *ReservationService) GetByID(id uint) (*models.Reservation, error) {
	return s.repo.GetByID(id)
}

// Get My Reservations
func (s *ReservationService) GetMyReservations(userID uint) ([]models.Reservation, error) {
	return s.repo.GetByUserID(userID)
}

// Cancel Reservation (Owner Only)
func (s *ReservationService) Cancel(userID uint, id uint) error {

	reservation, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Ownership Check
	if reservation.UserID != userID {
		return errors.New("you are not allowed to cancel this reservation")
	}

	reservation.Status = "cancelled"

	return s.repo.Update(reservation)
}

// Delete Reservation
func (s *ReservationService) Delete(id uint) error {
	return s.repo.Delete(id)
}
