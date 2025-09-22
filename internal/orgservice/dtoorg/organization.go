package dtoorg

import (
	"time"

	"github.com/google/uuid"
)

type UserMembershipsQuery struct {
	ID          uuid.UUID  `json:"iD" db:"id"`
	OrgID       uuid.UUID  `json:"orgID" db:"org_id"`
	RoleID      uuid.UUID  `json:"roleID" db:"role_id"`
	Status      string     `json:"status" db:"status"`
	FinalizedAt *time.Time `json:"finalizedAt" db:"finalized_at"`
	Name        string     `json:"name" db:"name"`
	TypeOrg     string     `json:"typeOrg" db:"type"`
}

type OrganizationMembershipsQuery struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	UserID      uuid.UUID  `json:"userId" db:"user_id"`
	RoleID      uuid.UUID  `json:"roleId" db:"role_id"`
	Status      string     `json:"status" db:"status"`
	FinalizedAt *time.Time `json:"finalizedAt" db:"finalized_at"`
	UserName    string     `json:"userName" db:"user_name"`
}
