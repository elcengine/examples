package dto

import (
	"mailman/src/modules/users/api/v1/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserReq struct {
	Name          string `validate:"required"`
	Email         string `validate:"required,email"`
	Role          string
	Organizations []string
}

type CreateUserRes struct {
	ID       primitive.ObjectID `json:"_id"`
	Password string             `json:"password"`
}

type GetAllUsersRes = []models.User

type GetUserByIDRes = models.User

type UpdateUserReq struct {
	Name          string
	Email         string `validate:"email"`
	Role          string
	Organizations []string
}

type UpdateUserRes = models.User

type DeleteUserRes = models.User
