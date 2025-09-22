package serviceauth

import (
	"proyecto/internal/authservice/storeauth"
	"proyecto/internal/shared/mailer"
)

type ServiceAuth struct {
	store  storeauth.StoreAuth
	mailer mailer.ResendMailer
}

func NewAuthService(s storeauth.StoreAuth, mailer mailer.ResendMailer) *ServiceAuth {
	return &ServiceAuth{store: s, mailer: mailer}
}
