package orgservice

import (
	"database/sql"
	"proyecto/internal/authservice/transport"
	"proyecto/internal/orgservice/serviceorg"
	"proyecto/internal/orgservice/storeorg"
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/tokenizer"
)

func RunOrg(db *sql.DB, tokenizerJWT tokenizer.TokenizerJWT, mailersmtp mailer.Mailer, authService *service.ServiceAuth) (*service.ServiceAuth, *transport.Handler) {

	orgStore := storeorg.NewOrgStore(db)
	orgService := serviceorg.NewOrgService(orgStore, tokenizerJWT, mailersmtp)
	orgHandler := transport.NewOrgHandler(orgService)

	orgHandler.SetupOrgRoutes()

	return orgService, orgHandler

}
