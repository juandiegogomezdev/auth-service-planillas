package transportorg

import (
	"net/http"
	"proyecto/internal/orgservice/serviceorg"
)

func (h *HandlerOrg) SetupOrgRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/session", h.HandlerSessionOrg)
	mux.HandleFunc("/permission", h.HandlerPermissionOrg)
	// mux.Handle("/org", middleware.AuthMiddleware(http.HandlerFunc(h.HandlerOrganization)))
	// mux.HandleFunc("/org/permission", middleware.AuthMiddleware(http.HandlerFunc(h.HandlerOrgPermission)))

}

type HandlerOrg struct {
	service *serviceorg.ServiceOrg
}

func NewOrgHandler(s *serviceorg.ServiceOrg) *HandlerOrg {
	return &HandlerOrg{service: s}
}
