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

type UserWithPasswordHash struct {
	User         models.UserModel
	PasswordHash string
}

func (ur *UsersRepository) GetWithPassword(email string) (UserWithPasswordHash, error) {
	var user UserWithPasswordHash

	err := ur.DbPool.QueryRow(
		context.Background(),
		"SELECT u.id, u.name, u.social_name, u.birthday, up.password FROM users u INNER JOIN user_passwords up ON up.user_id = u.id WHERE u.email = $1",
		email,
	).Scan(&user.User.Id, &user.User.Name, &user.User.SocialName, &user.User.Birthday, &user.PasswordHash)

	if err != nil {
		return UserWithPasswordHash{}, err
	}

	user.User.Email = email

	return user, nil
}
