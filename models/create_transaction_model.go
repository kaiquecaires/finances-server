package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateTransactionModel struct {
	Description    string          `json:"description" validate:"required"`
	BillCategoryId string          `json:"bill_category_id" validate:"required"`
	Amount         decimal.Decimal `json:"amount" validate:"required"`
	Date           time.Time       `json:"date" validate:"required"`
	Type           string          `json:"type" validate:"required,oneof=income outcome"`
	UserId         string          `json:"user_id"`
}
