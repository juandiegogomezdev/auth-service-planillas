package modelorg

import "github.com/google/uuid"

type Roles struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

type Permissions struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

type RolePermissions struct {
	RoleID       uuid.UUID `json:"roleId" db:"role_id"`
	PermissionID uuid.UUID `json:"permissionId" db:"permission_id"`
}

// type OrgRolePermissions struct {
// 	OrgID        uuid.UUID `json:"org_id"`
// 	RoleID       uuid.UUID `json:"role_id"`
// 	PermissionID uuid.UUID `json:"permission_id"`
// 	Allowed      bool      `json:"allowed"`
// }
