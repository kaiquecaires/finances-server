package models

import "time"

type SignupModel struct {
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required,min=8,max=64"`
	Name       string    `json:"name" validate:"required"`
	SocialName string    `json:"social_name" validate:"required"`
	Birthday   time.Time `json:"birthday" validate:"required"`
}
