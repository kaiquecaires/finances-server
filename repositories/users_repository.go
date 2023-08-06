package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/models"
)

type UsersRepository struct {
	DbPool *pgxpool.Pool
}

func (ur *UsersRepository) Create(data models.SignupModel) (models.UserModel, error) {
	var insertedID string
	err := ur.DbPool.QueryRow(
		context.Background(),
		"INSERT INTO users (email, name, social_name, birthday) VALUES ($1, $2, $3, $4) RETURNING id",
		data.Email,
		data.Name,
		data.SocialName,
		data.Birthday,
	).Scan(&insertedID)

	if err != nil {
		return models.UserModel{}, err
	}

	return models.UserModel{
		Id:         insertedID,
		Email:      data.Email,
		Name:       data.Name,
		SocialName: data.SocialName,
		Birthday:   data.Birthday,
	}, nil
}
