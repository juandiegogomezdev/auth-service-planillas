package service

import (
	"proyecto/authservice/internal/store"
	"proyecto/internal/mailer"
	"proyecto/internal/tokenizer"
)

type Service struct {
	store     store.Store
	tokenizer tokenizer.TokenizerJWT
	mailer    mailer.Mailer
}
