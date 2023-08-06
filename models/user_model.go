package models

type UserModel struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	SocialName string `json:"social_name"`
	Birthday   string `json:"birthday"`
}
