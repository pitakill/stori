package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Bank      string
	Number    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
