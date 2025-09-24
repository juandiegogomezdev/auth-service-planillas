package serviceauth

import (
	"fmt"
	"proyecto/internal/authservice/apperrors"
	"proyecto/internal/shared/tokenizer"
	"proyecto/internal/shared/utils"
)

func (s *ServiceAuth) Login(email string, password string) (string, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		return "", apperrors.WrapSerror("user_not_found", fmt.Errorf("error fetching user by email: %w", err))
	}

	// Compare hash with password
	err = utils.CompareHashAndPassword(user.HashedPassword, password)
	if err != nil {
		return "", apperrors.WrapSerror("password_incorrect", fmt.Errorf("error comparing password: %w", err))
	}

	// Generate a code and store it in the DB with expiration
	code := utils.GenerateCode(6)
	createdAt := utils.TimeNow()
	expiresAt := utils.TimeAddMinutes(createdAt, 15)

	err = s.store.UpdateCode(user.ID, code, createdAt, expiresAt)
	if err != nil {
		return "", apperrors.WrapSerror("code_storage", fmt.Errorf("error storing code: %w", err))
	}

	// Generate JWT token
	token, err := tokenizer.JWTGenerateConfirmLoginToken(user.Email)
	if err != nil {
		return "", apperrors.WrapSerror("token_generation", fmt.Errorf("error generating login token: %w", err))
	}

	// Send code via email
	go s.mailer.Send(user.Email, "Your login code", fmt.Sprintf("Your login code is: %s. It expires in 15 minutes.", code))

	return token, nil
}

func (s *ServiceAuth) ConfirmLoginCode(token string, code string) (string, error) {
	tokenPayload, err := tokenizer.JWTParseConfirmLoginToken(token)
	if err != nil {
		return "", apperrors.WrapSerror("token_parsing", fmt.Errorf("error parsing login token: %w", err))
	}

	fmt.Println("Token payload:", tokenPayload.Email)

	verificationInfo, err := s.store.GetVerificationByEmail(tokenPayload.Email)
	if err != nil {
		return "", apperrors.WrapSerror("verification_not_found", fmt.Errorf("error fetching verification info by email: %w", err))
	}

	// Check if code matches
	if code != verificationInfo.Code {
		return "", apperrors.WrapSerror("invalid_code", fmt.Errorf("invalid code"))
	}
	// Check if code is expired
	if utils.TimeNow().After(verificationInfo.ExpiresAt) {
		return "", apperrors.WrapSerror("code_expired", fmt.Errorf("code has expired"))
	}

	// Generate new JWT token
	newToken, err := tokenizer.JWTGenerateOrgSelectToken(verificationInfo.ID)
	if err != nil {
		return "", apperrors.WrapSerror("token_generation", fmt.Errorf("error generating new login token: %w", err))
	}

	return newToken, nil
}
