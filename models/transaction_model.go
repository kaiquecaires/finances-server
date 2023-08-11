package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionModel struct {
	Id             string          `json:"id"`
	Description    string          `json:"description"`
	BillCategoryId string          `json:"bil_category_id"`
	UserId         string          `json:"user_id"`
	Amount         decimal.Decimal `json:"amount"`
	Date           time.Time       `json:"date"`
	Type           string          `json:"transaction"`
}
