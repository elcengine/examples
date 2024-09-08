package middleware

import (
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"

	e_utils "github.com/elcengine/elemental/utils"
	"github.com/gofiber/fiber/v2"
)

func Protect(ctx *fiber.Ctx) error {
	token := ctx.Get(fiber.HeaderAuthorization)
	if token == "" {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Missing bearer token"))
	}
	user := utils.ValidateUserJWTToken(token[len("Bearer "):])
	ctx.Locals("user", user)
	return ctx.Next()
}

func AdminProtect(ctx *fiber.Ctx) error {
	user := e_utils.Cast[(*models.User)](ctx.Locals("user"))
	if user == nil || user.Role != models.Admin {
		panic(fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource"))
	}
	return ctx.Next()
}
