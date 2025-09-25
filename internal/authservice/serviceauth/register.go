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
	tokenConfirmRegister, err := tokenizer.JWTGenerateConfirmRegisterToken(email)
	if err != nil {
		return fmt.Errorf("error generating confirmation token: %w", err)
	}

	// Send confirmation email
	go s.mailer.Send(email, "Confirm your email", "Confirm your email: <a href='"+config.STATIC_CONFIRM_EMAIL_URL+"?token="+tokenConfirmRegister+"'>Click here</a>")

	return nil
}

func (s *ServiceAuth) CreateUser(token string, password string) (string, error) {
	claims, err := tokenizer.JWTParseConfirmRegisterToken(token)
	if err != nil {
		return "", apperrors.WrapSerror("invalid_token", fmt.Errorf("invalid or expired token: %w", err))
	}

	email := claims.NewEmail

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	code := utils.GenerateCode(6)
	createdAt := utils.TimeNow()
	expiresAt := createdAt.Add(1 * time.Minute)

	exist, err := s.store.ExistUser(email)
	if err != nil {
		return "", fmt.Errorf("error checking if user exists: %w", err)
	}
	if exist {
		return "", apperrors.WrapSerror("user_exists", fmt.Errorf("user already exists"))
	}

	userUUID := uuid.New()
	err = s.store.CreateUser(&modelauth.User{
		ID:             userUUID,
		Email:          email,
		HashedPassword: string(hashedPassword),
		CreatedAt:      createdAt,
	}, code, createdAt, expiresAt)

	if err != nil {
		return "", apperrors.WrapSerror("user_creation", fmt.Errorf("error creating user: %w", err))
	}

	// Generate token for org selection page
	tokenOrgSelect, err := tokenizer.JWTGenerateOrgSelectToken(userUUID)
	if err != nil {
		return "", apperrors.WrapSerror("token_generation", fmt.Errorf("error generating org-select token: %w", err))
	}

	return tokenOrgSelect, nil
}
