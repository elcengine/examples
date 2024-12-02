package system

import (
	"mailman/src/global"
	"mailman/src/modules/system/dto"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.JSON(global.Response[*interface{}]{
		Message: "Mailman up and running",
	})
}

func Memory(c *fiber.Ctx) error {
	return c.JSON(global.Response[dto.MemStats]{
		Message: "Memory usage retrieved",
		Data:    GetMemoryUsage(),
	})
}

//endpoint to return okay after the given delay in milliseconds
func SimulateCompute(c *fiber.Ctx) error {
	delayStr := c.Params("compute")
	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.Response[*interface{}]{
			Message: "Invalid compute parameter",
		})
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	return c.JSON(global.Response[*interface{}]{
		Message: "compute completed",
	})
}
