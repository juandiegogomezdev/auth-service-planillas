package storeorg

import (
	"fmt"
	"proyecto/internal/orgservice/dtoorg"
	"proyecto/internal/orgservice/modelorg"

	"github.com/google/uuid"
)

func (s *store) ExistOrganizationPersonal(ownerID string) (bool, error) {
	qCheck := `SELECT EXISTS(SELECT 1 FROM organizations WHERE owner_user_id=$1 AND type=$2)`
	var exists bool
	if err := s.db.Get(&exists, qCheck, ownerID, "personal"); err != nil {
		return false, fmt.Errorf("error checking existing organization: %w", err)
	}
	return exists, nil
}

func (s *store) CreateNewOrganization(org modelorg.Organizations) error {
	// First check if an organization with the same type exists for the user

	qCreate := `INSERT INTO organizations (id, owner_user_id, name, type, created_at)
				VALUES (:id, :owner_user_id, :name, :type, :created_at)`

	if _, err := s.db.NamedExec(qCreate, org); err != nil {
		return fmt.Errorf("error creating organization in store: %w", err)
	}
	return nil
}

func (s *store) GetAllUserOrganizations(userID string) ([]dtoorg.CreatedOrganization, error) {
	var organizations []dtoorg.CreatedOrganization
	q := `SELECT id, name, type, created_at FROM organizations WHERE owner_user_id=$1`
	if err := s.db.Select(&organizations, q, userID); err != nil {
		return nil, fmt.Errorf("error fetching organizations: %w", err)
	}
	return organizations, nil
}

func (s *store) CreateUserMembership(membership *modelorg.OrganizationMemberships) error {
	q := `
			INSERT INTO organization_memberships (id, org_id, user_id, role_id, status, created_at, created_by)
			VALUES (:id, :org_id, :user_id, :role_id, :status, :created_at, :created_by)
		`
	if _, err := s.db.NamedExec(q, membership); err != nil {
		return fmt.Errorf("error creating user membership: %w", err)
	}
	return nil

}

// All memberships of a user
// This is necesary for the user to see all orgs they belong to even if they are revoked
func (s *store) GetUserMemberships(userID uuid.UUID) ([]dtoorg.UserMembershipsQuery, error) {
	q := `
			SELECT om.id, om.org_id, om.role_id, om.status, om.finalized_at, o.name
			FROM organization_memberships om
			JOIN organizations o ON om.org_id = o.id
			WHERE om.user_id=$1
		`
	var memberships []dtoorg.UserMembershipsQuery
	if err := s.db.Select(&memberships, q, userID); err != nil {
		return nil, fmt.Errorf("error fetching user memberships: %w", err)
	}
	return memberships, nil
}

// All memberships of an organization
// This is necesary for the org admin to see all members/exmembers of the org
func (s *store) GetOrganizationMemberships(orgID uuid.UUID) ([]dtoorg.OrganizationMembershipsQuery, error) {
	q := `
	SELECT om.id, om.user_id, om.role_id, om.status, om.finalized_at, u.name AS user_name
	FROM organization_memberships om
	JOIN users u ON om.user_id = u.id
	WHERE om.org_id=$1
	`

	var memberships []dtoorg.OrganizationMembershipsQuery
	if err := s.db.Select(&memberships, q, orgID); err != nil {
		return nil, fmt.Errorf("error fetching organization memberships: %w", err)
	}
	return memberships, nil
}

//
