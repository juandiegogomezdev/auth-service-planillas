package tokenizer

import (
	"fmt"
	"proyecto/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	TokenTypeConfirmRegister TokenType = "confirm_register"
	TokenTypeConfirmLogin    TokenType = "confirm_login"
	TokenTypeOrgSelect       TokenType = "org_select"
	TokenTypeAppAccess       TokenType = "app_access"
	TokenTypeUnkown          TokenType = "unknown"
)

type BaseClaims struct {
	TokenType string `json:"type"`
	jwt.RegisteredClaims
}

// Used for ConfirmEmailHandler
type ConfirmRegisterClaims struct {
	NewEmail string `json:"new_email"`
	BaseClaims
}

// Used for ConfirmLoginHandler
type ConfirmLoginClaims struct {
	Email string `json:"email"`
	BaseClaims
}

// Used for orgSelectHandler
type orgSelectClaims struct {
	UserUUID string `json:"user_uuid"`
	BaseClaims
}

// used for access to general endpoints
type AppAccessClaims struct {
	UserUUID       string `json:"user_uuid"`
	MembershipUUID string `json:"membership_uuid"`
	BaseClaims
}

// --- Identify token type ---
func IdentifyTokenType(tokenString string) (TokenType, error) {
	parsedClaims, err := parseToken(tokenString, &BaseClaims{})
	if err != nil {
		return "", err
	}

	claims := parsedClaims.(*BaseClaims)

	return TokenType(claims.TokenType), nil
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
	return token.Claims, nil

}

// --- Generators

// Generate a token for confirm register
func JWTGenerateConfirmRegisterToken(newEmail string) (string, error) {
	claims := ConfirmRegisterClaims{
		NewEmail: newEmail,
		BaseClaims: BaseClaims{
			TokenType: string(TokenTypeConfirmRegister),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			},
		},
	}

	return signToken(claims)
}

// Generate a token for confirm login
func JWTGenerateConfirmLoginToken(email string) (string, error) {
	claims := ConfirmLoginClaims{
		Email: email,
		BaseClaims: BaseClaims{
			TokenType: string(TokenTypeConfirmLogin),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		},
	}

	return signToken(claims)
}

// Generate a token for org selection
func JWTGenerateOrgSelectToken(userUUID uuid.UUID) (string, error) {
	claims := orgSelectClaims{
		UserUUID: userUUID.String(),
		BaseClaims: BaseClaims{
			TokenType: string(TokenTypeOrgSelect),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		},
	}

	return signToken(claims)
}

// Generate a token for app access
func JWTGenerateAppAccessToken(userUUID uuid.UUID, membershipUUID uuid.UUID) (string, error) {
	claims := AppAccessClaims{
		UserUUID:       userUUID.String(),
		MembershipUUID: membershipUUID.String(),
		BaseClaims: BaseClaims{
			TokenType: string(TokenTypeAppAccess),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		},
	}

	return signToken(claims)
}

// --- Parsers ---

// Parse the token for confirm register
func JWTParseConfirmRegisterToken(tokenString string) (*ConfirmRegisterClaims, error) {
	parsedClaims, err := parseToken(tokenString, &ConfirmRegisterClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*ConfirmRegisterClaims), nil
}

// Parse the token for confirm login
func JWTParseConfirmLoginToken(tokenString string) (*ConfirmLoginClaims, error) {
	parsedClaims, err := parseToken(tokenString, &ConfirmLoginClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*ConfirmLoginClaims), nil
}

// Parse the token for org selection
func JWTParseOrgSelectToken(tokenString string) (*orgSelectClaims, error) {
	parsedClaims, err := parseToken(tokenString, &orgSelectClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*orgSelectClaims), nil
}

// Parse the token for app access
func JWTParseMembershipAccessToken(tokenString string) (*AppAccessClaims, error) {
	parsedClaims, err := parseToken(tokenString, &AppAccessClaims{})
	if err != nil {
		return nil, err
	}

	return parsedClaims.(*AppAccessClaims), nil
}
