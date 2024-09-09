package dto

import "mailman/src/modules/users/api/v1/models"

type LoginReq struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type LoginRes struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

type RegisterReq struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	Age      int    `validate:"required"`
}

type CreateUserDTO struct {
	Name string `augmented_validate:"unique=users" json:"name"`
	Email    string `validate:"required,email"`
	Age  int    `validate:"max=150,min=18" json:"age"`
	Password string `validate:"required,min=8"`
}
