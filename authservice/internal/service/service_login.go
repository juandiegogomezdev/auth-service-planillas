package service

import (
	"fmt"

	"proyecto/authservice/internal/store"
	"proyecto/internal/mailer"
	"proyecto/internal/tokenizer"
	"proyecto/internal/utils"
)

func NewAuthService(s store.Store, tokenizer tokenizer.TokenizerJWT, mailer mailer.Mailer) *Service {
	return &Service{store: s, tokenizer: tokenizer, mailer: mailer}
}

type LoginStatus int

const (
	LoginStatusInvalidCredentials LoginStatus = iota
	LoginStatusUserNotFound
	LoginStatusSuccess
	LoginPasswordIncorrect
)

func (s *Service) Login(email string, password string) (string, LoginStatus, error) {
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

func (s *Service) ConfirmLoginCode(token string, code string) (string, ConfirmLoginStatus, error) {
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
