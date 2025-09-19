package serviceauth

import (
	"proyecto/internal/authservice/storeauth"
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/tokenizer"
)

type ServiceAuth struct {
	store     storeauth.StoreAuth
	tokenizer tokenizer.TokenizerJWT
	mailer    mailer.Mailer
}

func NewAuthService(s storeauth.StoreAuth, tokenizer tokenizer.TokenizerJWT, mailer mailer.Mailer) *ServiceAuth {
	return &ServiceAuth{store: s, tokenizer: tokenizer, mailer: mailer}
}
