package utils

import (
	"math/rand"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateCode(n int) string {
	numbers := "0123456789"
	code := make([]byte, n)
	for i := range code {
		code[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(code)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
