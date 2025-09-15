package tokenizer

import (
	"fmt"
	"proyecto/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenizerJWT interface {
	GenerateConfirmEmailToken(newEmail string) (string, error)
	ParseConfirmEmailToken(tokenStr string) (*JWTConfirmEmailClaims, error)

	GenerateLoginToken(email string) (string, error)
	ParseLoginToken(tokenStr string) (*JWTLoginClaims, error)

	GenerateAccessToken(userUUID uuid.UUID) (string, error)
	ParseAccessToken(tokenStr string) (*JWTAccessClaims, error)
}

type tokenizerJWT struct{}

func NewTokenizerJWT() TokenizerJWT {
	return &tokenizerJWT{}
}

type JWTConfirmEmailClaims struct {
	NewEmail string
	jwt.RegisteredClaims
}

type JWTLoginClaims struct {
	Email string
	jwt.RegisteredClaims
}

type JWTAccessClaims struct {
	UserUUID string
	jwt.RegisteredClaims
}

func (t *tokenizerJWT) GenerateConfirmEmailToken(newEmail string) (string, error) {
	claims := JWTConfirmEmailClaims{
		NewEmail: newEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 200)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

func (t *tokenizerJWT) GenerateLoginToken(email string) (string, error) {
	claims := JWTLoginClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

func (t *tokenizerJWT) GenerateAccessToken(userUUID uuid.UUID) (string, error) {
	claims := JWTAccessClaims{
		UserUUID: userUUID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil

}

func (t *tokenizerJWT) ParseConfirmEmailToken(tokenString string) (*JWTConfirmEmailClaims, error) {
	claims := &JWTConfirmEmailClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func (t *tokenizerJWT) ParseLoginToken(tokenString string) (*JWTLoginClaims, error) {
	claims := &JWTLoginClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func (t *tokenizerJWT) ParseAccessToken(tokenString string) (*JWTAccessClaims, error) {
	claims := &JWTAccessClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// func (t *tokenizerJWT) GenerateAccessToken(userUUID string) (string, error) {
// 	claims := JWTAccessClaims{
// 		UserUUID: userUUID,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString(config.Config.JWTSecret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func (t *tokenizerJWT) ParseAccessToken(tokenString string) (*JWTAccessClaims, error) {
// 	claims := &JWTAccessClaims{}

// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return config.Config.JWTSecret, nil
// 	}, jwt.WithValidMethods([]string{"HS256"}))

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	return claims, nil
// }
