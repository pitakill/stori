package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Date      time.Time
	Credit    bool
	Amount    decimal.Decimal
	CreatedAt time.Time
	UpdatedAt time.Time
}
