package service

import (
	"errors"

	"github.com/Mokarama/assignment-6-spotsync-api/dto"
	"github.com/Mokarama/assignment-6-spotsync-api/models"
	"github.com/Mokarama/assignment-6-spotsync-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: repository.NewAuthRepository(),
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) error {

	// Check if email already exists
	existingUser, _ := s.repo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "driver",
	}

	return s.repo.CreateUser(&user)
}
