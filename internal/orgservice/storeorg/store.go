package storeorg

import (
	"proyecto/internal/orgservice/dtoorg"
	"proyecto/internal/orgservice/modelorg"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StoreOrg interface {
	// Organizations
	GetAllUserOrganizations(userID string) ([]dtoorg.CreatedOrganization, error)
	GetUserMemberships(userID uuid.UUID) ([]dtoorg.UserMembershipsQuery, error)
	CreateUserMembership(membership *modelorg.OrganizationMemberships) error

	ExistOrganizationPersonal(ownerID string) (bool, error)
	GetOrganizationMemberships(orgID uuid.UUID) ([]dtoorg.OrganizationMembershipsQuery, error)
	CreateNewOrganization(org modelorg.Organizations) error

	// permissions
	GetAllRolePermissions() (map[uuid.UUID][]uuid.UUID, error)
	GetAllPermissions() ([]modelorg.Permissions, error)
	GetAllRoles() ([]modelorg.Roles, error)
	GetUserOrganizationRole(id uuid.UUID) (uuid.UUID, error)

	// No implemented yet
	GetRoleByMembership(membershipID uuid.UUID) (uuid.UUID, error)
}

type store struct {
	db *sqlx.DB
}

func NewOrgStore(db *sqlx.DB) StoreOrg {
	return &store{db: db}
}
