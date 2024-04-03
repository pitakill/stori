package domain

import "github.com/google/uuid"

type Email struct {
	AccountID uuid.UUID
	Summary   *Summary
}
