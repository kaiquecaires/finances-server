package models

type DeleteTransactionModel struct {
	Id string `json:"id" validate:"required"`
}
