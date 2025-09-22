package tokenizer

import (
	"fmt"
	"proyecto/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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

type JWTMembershipAccessClaims struct {
	UserUUID       string
	MembershipUUID string
	jwt.RegisteredClaims
}

// --- helpers ---
func signToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.JWT.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

func parseToken(tokenString string, claims jwt.Claims) (jwt.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.JWTSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil

}

// --- Generators

func JWTGenerateConfirmEmailToken(newEmail string) (string, error) {
	claims := JWTConfirmEmailClaims{
		NewEmail: newEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 200)),
		},
	}

	return signToken(claims)
}

func JWTGenerateLoginToken(email string) (string, error) {
	claims := JWTLoginClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	return signToken(claims)
}

func JWTGenerateAccessToken(userUUID uuid.UUID) (string, error) {
	claims := JWTAccessClaims{
		UserUUID: userUUID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	return signToken(claims)
}

func JWTGenerateMembershipAccessToken(userUUID uuid.UUID, membershipUUID uuid.UUID) (string, error) {
	claims := JWTMembershipAccessClaims{
		UserUUID:       userUUID.String(),
		MembershipUUID: membershipUUID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	return signToken(claims)
}

// --- Parsers ---

func JWTParseConfirmEmailToken(tokenString string) (*JWTConfirmEmailClaims, error) {
	parsedClaims, err := parseToken(tokenString, &JWTConfirmEmailClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*JWTConfirmEmailClaims), nil
}

func JWTParseLoginToken(tokenString string) (*JWTLoginClaims, error) {
	parsedClaims, err := parseToken(tokenString, &JWTLoginClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*JWTLoginClaims), nil
}

func JWTParseAccessToken(tokenString string) (*JWTAccessClaims, error) {
	parsedClaims, err := parseToken(tokenString, &JWTAccessClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*JWTAccessClaims), nil
}

func JWTParseMembershipAccessToken(tokenString string) (*JWTMembershipAccessClaims, error) {
	parsedClaims, err := parseToken(tokenString, &JWTMembershipAccessClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*JWTMembershipAccessClaims), nil
}
