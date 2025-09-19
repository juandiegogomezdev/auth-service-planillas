package serviceorg

import (
	"proyecto/internal/orgservice/storeorg"
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/tokenizer"
)

type ServiceOrg struct {
	store     storeorg.StoreOrg
	tokenizer tokenizer.TokenizerJWT
	mailer    mailer.Mailer
}

func NewOrgService(s storeorg.StoreOrg, tokenizer tokenizer.TokenizerJWT, mailer mailer.Mailer) *ServiceOrg {
	return &ServiceOrg{store: s, tokenizer: tokenizer, mailer: mailer}
}
