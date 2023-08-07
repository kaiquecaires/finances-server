package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/helpers"
)

type UserPasswordsRepository struct {
	DbPool *pgxpool.Pool
}

func (up *UserPasswordsRepository) Create(userId string, password string) (string, error) {
	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		return "", err
	}

	var insertedId string

	err = up.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO user_passwords (user_id, password) VALUES ($1, $2) RETURNING id",
		userId,
		hashedPassword,
	).Scan(&insertedId)

	return insertedId, err
}
