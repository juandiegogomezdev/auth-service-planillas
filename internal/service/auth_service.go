package service

import (
	"proyecto/internal/model"
	"proyecto/internal/store"
	"proyecto/internal/utils"

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

	createdAt := utils.TimeNow()

	return s.store.CreateUser(&model.User{
		ID:             uuid.New(),
		Email:          email,
		HashedPassword: string(hashedPassword),
		CreatedAt:      createdAt,
	})
}
