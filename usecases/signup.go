package usecases

import "github.com/kaiquecaires/finances-server/models"

func Signup(s *models.SignupModel) models.UserModel {
	return models.UserModel{
		Id:         "any",
		Email:      s.Email,
		Name:       s.Name,
		SocialName: s.SocialName,
		Birthday:   s.Birthday,
	}
}
