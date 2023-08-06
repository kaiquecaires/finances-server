package models

type SignupModel struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8,max=64"`
	Name       string `json:"name" validate:"required"`
	SocialName string `json:"socialName" validate:"required"`
	Birthday   string `json:"birthday" validate:"required"`
}
