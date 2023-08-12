package models

type ListTransactionsModel struct {
	Limit int `form:"limit" validate:"required"`
	Page  int `form:"page" validate:"required"`
}
