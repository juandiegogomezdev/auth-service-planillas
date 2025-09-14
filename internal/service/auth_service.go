package service

import (
	"proyecto/internal/model"
	"proyecto/internal/store"
	"proyecto/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	store store.AuthStore
}

func NewAuthService(s store.AuthStore) *AuthService {
	return &AuthService{store: s}
}

func (s *AuthService) CreateUser(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	code := utils.GenerateCode(6)
	createdAt := utils.TimeNow()
	expiresAt := createdAt.Add(1 * time.Minute)

	return s.store.CreateUser(&model.User{
		ID:             uuid.New(),
		Email:          email,
		HashedPassword: string(hashedPassword),
		CreatedAt:      createdAt,
	}, code, createdAt, expiresAt)
}

func (s *AuthService) Login(email string, password string) error {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		return err
	}

	// Compare hash with password
	err = utils.CompareHashAndPassword(user.HashedPassword, password)
	if err != nil {
		return err
	}

	//

	return nil
}
