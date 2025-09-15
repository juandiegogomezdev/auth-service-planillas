package service

import (
	"fmt"
	"log"
	"proyecto/internal/mailer"
	"proyecto/internal/model"
	"proyecto/internal/store"
	"proyecto/internal/tokenizer"
	"proyecto/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	store     store.AuthStore
	tokenizer tokenizer.TokenizerJWT
	mailer    mailer.Mailer
}

func NewAuthService(s store.AuthStore, tokenizer tokenizer.TokenizerJWT, mailer mailer.Mailer) *AuthService {
	return &AuthService{store: s, tokenizer: tokenizer, mailer: mailer}
}

type RegisterUserStatus int

const (
	RegisterUserStatusUserExists RegisterUserStatus = iota
	RegisterUserStatusConfirmationSent
)

func (s *AuthService) RegisterUser(email string) (RegisterUserStatus, error) {
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

func (s *AuthService) CreateUser(token string, password string) (CreateUserStatus, error) {
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

type LoginStatus int

const (
	LoginStatusInvalidCredentials LoginStatus = iota
	LoginStatusUserNotFound
	LoginStatusSuccess
	LoginPasswordIncorrect
)

func (s *AuthService) Login(email string, password string) (string, LoginStatus, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		return "", LoginStatusUserNotFound, fmt.Errorf("error fetching user by email: %w", err)
	}
	// Compare hash with password
	err = utils.CompareHashAndPassword(user.HashedPassword, password)
	if err != nil {
		return "", LoginPasswordIncorrect, fmt.Errorf("error comparing password: %w", err)
	}

	// Generate JWT token
	token, err := s.tokenizer.GenerateLoginToken(user.Email)
	if err != nil {
		return "", LoginStatusSuccess, fmt.Errorf("error generating login token: %w", err)
	}

	return token, LoginStatusSuccess, nil
}

type ConfirmLoginStatus int

const (
	ConfirmLoginStatusInvalidToken ConfirmLoginStatus = iota
	ConfirmLoginStatusSuccess
	ConfirmLoginInvalidCode
)

func (s *AuthService) ConfirmLoginCode(token string, code string) (string, ConfirmLoginStatus, error) {
	tokenPayload, err := s.tokenizer.ParseLoginToken(token)
	if err != nil {
		return "", ConfirmLoginStatusInvalidToken, fmt.Errorf("error parsing login token: %w", err)
	}

	verificationInfo, err := s.store.GetVerificationByEmail(tokenPayload.Email)
	if err != nil {
		return "", ConfirmLoginStatusInvalidToken, fmt.Errorf("error fetching code by email: %w", err)
	}

	if code != verificationInfo.Code {
		return "", ConfirmLoginInvalidCode, fmt.Errorf("invalid code")
	}

	// Generate new JWT token
	newToken, err := s.tokenizer.GenerateAccessToken(verificationInfo.ID)
	if err != nil {
		return "", ConfirmLoginStatusInvalidToken, fmt.Errorf("error generating new login token: %w", err)
	}

	return newToken, ConfirmLoginStatusSuccess, nil
}
