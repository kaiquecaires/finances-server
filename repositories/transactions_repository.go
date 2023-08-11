package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/models"
)

type TransactionsRepository struct {
	DbPool *pgxpool.Pool
}

func (t TransactionsRepository) Create(data models.CreateTransactionModel) (models.TransactionModel, error) {
	var insertedId string

	err := t.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO transactions (description, type, bill_category_id, user_id, amount, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		data.Description,
		data.Type,
		data.BillCategoryId,
		data.UserId,
		data.Amount,
		data.Date,
	).Scan(&insertedId)

	if err != nil {
		return models.TransactionModel{}, err
	}

	return models.TransactionModel{
		Id:             insertedId,
		Description:    data.Description,
		BillCategoryId: data.BillCategoryId,
		UserId:         data.UserId,
		Amount:         data.Amount,
		Date:           data.Date,
		Type:           data.Type,
	}, nil
}
