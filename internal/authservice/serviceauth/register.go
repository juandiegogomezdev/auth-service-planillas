package serviceauth

import (
	"fmt"
	"proyecto/config"
	"proyecto/internal/authservice/apperrors"
	"proyecto/internal/authservice/modelauth"
	"proyecto/internal/shared/tokenizer"
	"proyecto/internal/shared/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceAuth) RegisterUser(email string) error {
	exists, err := s.store.ExistUser(email)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}
	if exists {
		return apperrors.WrapSerror("user_exists", fmt.Errorf("user already exists"))
	}

	// Generate a token with the email
	token, err := tokenizer.JWTGenerateConfirmEmailToken(email)
	if err != nil {
		return fmt.Errorf("error generating confirmation token: %w", err)
	}

	// Send confirmation email
	go s.mailer.Send(email, "Confirm your email", "Confirm your email: <a href='"+config.STATIC_CONFIRM_EMAIL_URL+"?token="+token+"'>Click here</a>")

	return nil
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

	fmt.Println("Creating user with email:", email)
	fmt.Println("Password:", password)

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
