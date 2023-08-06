package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordsRepository struct {
	DbPool *pgxpool.Pool
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (up *UserPasswordsRepository) Create(userId string, password string) (string, error) {
	hashedPassword, err := hashPassword(password)

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
