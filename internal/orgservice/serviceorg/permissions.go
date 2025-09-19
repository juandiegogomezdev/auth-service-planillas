package serviceorg

import (
	"fmt"
	"proyecto/internal/orgservice/dtoorg"
)

// PermissionsInfo retrieves all roles, permissions, and their associations.
// This should be cached on the service layer to avoid frequent DB hits.
func (s *ServiceOrg) PermissionsInfo() (dtoorg.PermissionsInfoResponse, error) {
	roles, err := s.store.GetAllRoles()
	if err != nil {
		return dtoorg.PermissionsInfoResponse{}, fmt.Errorf("error fetching roles: %w", err)
	}
	permissions, err := s.store.GetAllPermissions()
	if err != nil {
		return dtoorg.PermissionsInfoResponse{}, fmt.Errorf("error fetching permissions: %w", err)
	}

	rolPermissions, err := s.store.GetAllRolePermissions()
	if err != nil {
		return dtoorg.PermissionsInfoResponse{}, fmt.Errorf("error fetching role-permissions mapping: %w", err)
	}

	permissionsInfo := dtoorg.PermissionsInfoResponse{
		Roles:             roles,
		Permissions:       permissions,
		PermissionsByRole: rolPermissions,
	}

	return permissionsInfo, nil

}
