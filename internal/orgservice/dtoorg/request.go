package dtoorg

import (
	"time"

	"github.com/google/uuid"
)

type CreatedOrganization struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// type AlterPermissionRequest struct {
// 	RoleID       uuid.UUID `json:"role_id"`
// 	PermissionID uuid.UUID `json:"permission_id"`
// 	Allowed      bool      `json:"allowed"`
// }
