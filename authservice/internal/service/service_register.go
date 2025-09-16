package service

import (
	"fmt"
	"log"
	"proyecto/authservice/internal/model"
	"proyecto/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserStatus int

const (
	RegisterUserStatusUserExists RegisterUserStatus = iota
	RegisterUserStatusConfirmationSent
)

func (s *Service) RegisterUser(email string) (RegisterUserStatus, error) {
	exists, err := s.store.ExistUser(email)
	if err != nil {
		return 0, err
	}
	if exists {
		return RegisterUserStatusUserExists, nil
	}

	token, err := s.tokenizer.GenerateConfirmEmailToken(email)
	if err != nil {
		return 0, err
	}
	fmt.Println("Generated token:", token)

	// Send confirmation email
	err = s.mailer.Send(email, "Confirm your email", "your token is: "+token)
	if err != nil {
		log.Println("Error sending email:", err)
		return 0, err
	}

	return RegisterUserStatusConfirmationSent, nil
}

type CreateUserStatus int

const (
	CreateUserStatusUserExists CreateUserStatus = iota
	CreateUserStatusUserCreated
)

func (s *Service) CreateUser(token string, password string) (CreateUserStatus, error) {
	claims, err := s.tokenizer.ParseConfirmEmailToken(token)
	if err != nil {
		return 0, err
	}

	email := claims.NewEmail

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	code := utils.GenerateCode(6)
	createdAt := utils.TimeNow()
	expiresAt := createdAt.Add(1 * time.Minute)

	exist, err := s.store.ExistUser(email)
	if err != nil {
		return 0, err
	}
	if exist {
		return CreateUserStatusUserExists, nil
	}

	err = s.store.CreateUser(&model.User{
		ID:             uuid.New(),
		Email:          email,
		HashedPassword: string(hashedPassword),
		CreatedAt:      createdAt,
	}, code, createdAt, expiresAt)

	if err != nil {
		return 0, err
	}

	return CreateUserStatusUserCreated, nil
}
