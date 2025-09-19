package modelorg

import (
	"time"

	"github.com/google/uuid"
)

type Organizations struct {
	ID        uuid.UUID `json:"id" db:"id"`
	OwnerID   uuid.UUID `json:"ownerUserId" db:"owner_user_id"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type OrganizationMemberships struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	OrgID       uuid.UUID  `json:"orgId" db:"org_id"`
	UserID      uuid.UUID  `json:"userId" db:"user_id"`
	RoleID      uuid.UUID  `json:"roleId" db:"role_id"`
	Status      string     `json:"status" db:"status"`
	CreatedBy   uuid.UUID  `json:"createdBy" db:"created_by"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	FinalizedBy *uuid.UUID `json:"finalizedBy" db:"finalized_by"`
	FinalizedAt *time.Time `json:"finalizedAt" db:"finalized_at"`
}
