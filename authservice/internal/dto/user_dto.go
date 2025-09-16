package dto

import (
	"time"

	"github.com/google/uuid"
)

type UsersQuery struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
}
