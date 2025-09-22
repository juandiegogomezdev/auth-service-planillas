package orgservice

import (
	"net/http"
	"proyecto/internal/orgservice/serviceorg"
	"proyecto/internal/orgservice/storeorg"
	"proyecto/internal/orgservice/transportorg"
	"proyecto/internal/shared/mailer"

	"github.com/jmoiron/sqlx"
)

func RunOrg(db *sqlx.DB, mailersmtp mailer.ResendMailer, mux *http.ServeMux) {

	orgStore := storeorg.NewOrgStore(db)
	orgService := serviceorg.NewOrgService(orgStore, mailersmtp)
	orgHandler := transportorg.NewOrgHandler(orgService)

	orgHandler.SetupOrgRoutes(mux)
}
