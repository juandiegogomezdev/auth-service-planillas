package service

import (
	"proyecto/internal/mailer"
	"proyecto/internal/tokenizer"
	"proyecto/services/authservice/internal/store"
)

type Service struct {
	store     store.Store
	tokenizer tokenizer.TokenizerJWT
	mailer    mailer.Mailer
}
