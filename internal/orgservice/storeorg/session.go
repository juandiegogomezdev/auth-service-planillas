package storeorg

import (
	"fmt"
	"proyecto/internal/orgservice/dtoorg"

	"github.com/google/uuid"
)

// All memberships of a user
// This is necesary for the user to see all orgs they belong to even if they are revoked
func (s *store) GetUserMemberships(userID uuid.UUID) ([]dtoorg.UserMembershipsQuery, error) {
	q := `
			SELECT om.id, om.org_id, om.role_id, om.status, om.finalized_at, o.name, o.type
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

// Get the membership ID of a user in a specific organization
func (s *store) CheckUserMembership(userID uuid.UUID, membershipID uuid.UUID) (uuid.UUID, error) {
	var memID uuid.UUID
	q := `SELECT id FROM organization_memberships WHERE user_id=$1 AND id=$2 AND status='active'`
	if err := s.db.Get(&memID, q, userID, membershipID); err != nil {
		return uuid.UUID{}, err
	}
	return memID, nil
}
