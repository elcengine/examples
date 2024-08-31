package v1

import (
	m "mailman/src/middleware"
	"mailman/src/modules/users/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	models.UserModel.SyncIndexes()
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateUserReq](m.Body), Create)
	return v1
}
