package v1

import (
	filter_query_middleware "github.com/elcengine/elemental/plugins/filter-query/middleware"
	"github.com/gofiber/fiber/v2"
	m "mailman/src/middleware"
	"mailman/src/modules/users/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"
)

func New() *fiber.App {
	models.UserModel.SyncIndexes()
	v1 := fiber.New()
	v1.Post("/", m.Validate[dto.CreateUserReq](m.Body), Create)
	v1.Get("/", filter_query_middleware.Parse, GetAll)
	v1.Get("/:id", GetUserByID)
	v1.Patch("/:id", m.Validate[dto.CreateUserReq](m.Body), Update)
	v1.Delete("/:id", Delete)
	return v1
}
