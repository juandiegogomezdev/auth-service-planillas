package serviceorg

import (
	"proyecto/internal/orgservice/storeorg"
	"proyecto/internal/shared/mailer"
)

type ServiceOrg struct {
	store  storeorg.StoreOrg
	mailer mailer.ResendMailer
}

func NewOrgService(s storeorg.StoreOrg, mailer mailer.ResendMailer) *ServiceOrg {
	return &ServiceOrg{store: s, mailer: mailer}
}
