package v1

import (
	"github.com/gofiber/fiber/v2"
	"mailman/src/global"
	"mailman/src/modules/users/api/v1/dto"
)

func Create(c *fiber.Ctx) error {
	payload := new(dto.CreateUserReq)
	c.BodyParser(payload)
	res := createUser(c, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.CreateUserRes]{
		Message: "User created successfully",
		Data:    &res,
	})
}

func GetAll(c *fiber.Ctx) error {
	res := getAllUsers(c)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.GetAllUsersRes]{
		Message: "Users fetched successfully",
		Data:    &res,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res := getUserByID(c, id)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.GetUserByIDRes]{
		Message: "User fetched successfully",
		Data:    &res,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	payload := new(dto.CreateUserReq)
	c.BodyParser(payload)
	res := updateUser(c, id, *payload)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.UpdateUserRes]{
		Message: "User updated successfully",
		Data:    &res,
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	res := deleteUser(c, id)
	return c.Status(fiber.StatusCreated).JSON(global.Response[dto.DeleteUserRes]{
		Message: "User deleted successfully",
		Data:    &res,
	})
}