package serviceauth

import (
	"log"
	"proyecto/config"
	"proyecto/internal/authservice/modelauth"
	"proyecto/internal/shared/tokenizer"
	"proyecto/internal/shared/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserStatus int

const (
	RegisterUserStatusUserExists RegisterUserStatus = iota
	RegisterUserStatusConfirmationSent
)

func (s *ServiceAuth) RegisterUser(email string) (RegisterUserStatus, error) {
	exists, err := s.store.ExistUser(email)
	if err != nil {
		return 0, err
	}
	if exists {
		return RegisterUserStatusUserExists, nil
	}

	token, err := tokenizer.JWTGenerateConfirmEmailToken(email)
	if err != nil {
		return 0, err
	}

	// Send confirmation email
	go s.mailer.Send(email, "Confirm your email", "Confirm your email: <a href='"+config.STATIC_CONFIRM_EMAIL_URL+"?token="+token+"'>Click here</a>")
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

func (s *ServiceAuth) CreateUser(token string, password string) (CreateUserStatus, error) {
	claims, err := tokenizer.JWTParseConfirmEmailToken(token)
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

	err = s.store.CreateUser(&modelauth.User{
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
