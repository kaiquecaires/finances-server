package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/models"
	"github.com/shopspring/decimal"
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

func (t TransactionsRepository) List(userId string, limit int, page int) ([]models.TransactionModel, error) {
	var transactions []models.TransactionModel
	offset := (page - 1) * limit

	rows, err := t.DbPool.Query(
		context.Background(),
		"SELECT id, description, bill_category_id, user_id, amount, date, type FROM transactions WHERE user_id = $1 AND deleted_at IS NULL ORDER BY date DESC LIMIT $2 OFFSET $3",
		userId,
		limit,
		offset,
	)

	if err != nil {
		return transactions, err
	}

	defer rows.Close()

	for rows.Next() {
		var transaction models.TransactionModel
		err := rows.Scan(
			&transaction.Id,
			&transaction.Description,
			&transaction.BillCategoryId,
			&transaction.UserId,
			&transaction.Amount,
			&transaction.Date,
			&transaction.Type,
		)
		if err != nil {
			break
		}
		transactions = append(transactions, transaction)
	}

	return transactions, err
}

func (t TransactionsRepository) GetAmount(userId string) (decimal.Decimal, error) {
	var amount decimal.Decimal

	err := t.DbPool.QueryRow(
		context.Background(),
		"SELECT SUM(amount) FROM transactions WHERE user_id = $1",
		userId,
	).Scan(&amount)

	return amount, err
}

func (t TransactionsRepository) GetTotal(userId string) (int, error) {
	var total int

	err := t.DbPool.QueryRow(
		context.Background(),
		"SELECT count(1) FROM transactions WHERE user_id = $1",
		userId,
	).Scan(&total)

	return total, err
}

func (t TransactionsRepository) Delete(id string) error {
	_, err := t.DbPool.Exec(
		context.Background(),
		"UPDATE transactions SET deleted_at = now() WHERE id = $1",
		id,
	)

	return err
}
