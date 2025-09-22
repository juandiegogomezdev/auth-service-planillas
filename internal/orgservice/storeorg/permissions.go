package storeorg

import (
	"proyecto/internal/orgservice/modelorg"

	"github.com/google/uuid"
)

// Get all roles in the system
func (s *store) GetAllRoles() ([]modelorg.Roles, error) {
	var roles []modelorg.Roles
	q := `SELECT id, name FROM roles`
	if err := s.db.Select(&roles, q); err != nil {
		return nil, err
	}
	return roles, nil
}

// Get all permissions in the system
func (s *store) GetAllPermissions() ([]modelorg.Permissions, error) {
	var permissions []modelorg.Permissions
	q := `SELECT id, name FROM permissions`
	if err := s.db.Select(&permissions, q); err != nil {
		return nil, err
	}
	return permissions, nil
}

// Get all permissions associated with each role
func (s *store) GetAllRolePermissions() (map[uuid.UUID][]uuid.UUID, error) {
	var rolePermissions []modelorg.RolePermissions

	q := `SELECT role_id, permission_id FROM role_permissions`

	if err := s.db.Select(&rolePermissions, q); err != nil {
		return nil, err
	}

	m := make(map[uuid.UUID][]uuid.UUID)
	for _, rp := range rolePermissions {
		m[rp.RoleID] = append(m[rp.RoleID], rp.PermissionID)
	}
	return m, nil
}

// Get the role ID of a user in a specific organization using the id membership
func (s *store) GetRoleByMembership(membershipID uuid.UUID) (uuid.UUID, error) {
	var roleID uuid.UUID
	q := `SELECT role_id FROM organization_memberships WHERE id=$1`
	if err := s.db.Get(&roleID, q, membershipID); err != nil {
		return uuid.UUID{}, err
	}
	return roleID, nil
}
