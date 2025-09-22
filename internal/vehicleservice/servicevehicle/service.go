package servicevehicle

import (
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/tokenizer"
	"proyecto/internal/vehicleservice/store"
)

type Service struct {
	store     store.Store
	tokenizer tokenizer.TokenizerJWT
	mailer    mailer.Mailer
}

func NewVehicleService(s store.Store, tokenizer tokenizer.TokenizerJWT, mailer mailer.Mailer) *Service {
	return &Service{store: s, tokenizer: tokenizer, mailer: mailer}
}
