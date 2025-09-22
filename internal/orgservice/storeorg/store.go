package storeorg

import (
	"proyecto/internal/orgservice/dtoorg"
	"proyecto/internal/orgservice/modelorg"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StoreOrg interface {
	// Session in org
	GetUserMemberships(userID uuid.UUID) ([]dtoorg.UserMembershipsQuery, error)
	CheckUserMembership(userID uuid.UUID, membershipID uuid.UUID) (uuid.UUID, error)

	// permissions
	GetAllRolePermissions() (map[uuid.UUID][]uuid.UUID, error)
	GetAllPermissions() ([]modelorg.Permissions, error)
	GetAllRoles() ([]modelorg.Roles, error)
	GetRoleByMembership(membershipID uuid.UUID) (uuid.UUID, error)

	// Organizations
	GetAllUserOrganizations(userID uuid.UUID) ([]dtoorg.CreatedOrganization, error)
	CreateUserMembership(membership *modelorg.OrganizationMemberships) error
	ExistOrganizationPersonal(ownerID uuid.UUID) (bool, error)
	GetOrganizationMemberships(orgID uuid.UUID) ([]dtoorg.OrganizationMembershipsQuery, error)
	CreateNewOrganization(org modelorg.Organizations) error

	//
}

type store struct {
	db *sqlx.DB
}

func NewOrgStore(db *sqlx.DB) StoreOrg {
	return &store{db: db}
}
