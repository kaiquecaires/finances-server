package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPasswordsRepository struct {
	DbPool *pgxpool.Pool
}

func (up *UserPasswordsRepository) Create(userId string, password string) error {
	err := up.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO user_passwords (userId, password) VALUES ($1, $2)",
		userId,
		password,
	).Scan()

	return err
}
