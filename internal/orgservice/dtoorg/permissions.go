package dtoorg

import (
	"proyecto/internal/orgservice/modelorg"

	"github.com/google/uuid"
)

type PermissionsInfoResponse struct {
	Roles             []modelorg.Roles          `json:"roles"`
	Permissions       []modelorg.Permissions    `json:"permissions"`
	PermissionsByRole map[uuid.UUID][]uuid.UUID `json:"permissionsByRole"`
}

type RolePermissions struct {
	RoleID        uuid.UUID   `json:"id"`
	IDPermissions []uuid.UUID `json:"permissions"`
}
