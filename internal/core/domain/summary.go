package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Summary struct {
	Total              decimal.Decimal
	TotalAverageDebit  decimal.Decimal
	TotalAverageCredit decimal.Decimal
	TransactionByMonth map[time.Month]int
}
