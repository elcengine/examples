package background

import (
	"mailman/src/modules/users/api/v1/models"

	"github.com/gofiber/fiber/v2/log"
)

func InitializeTasks() {
	models.UserModel.OnInsert(func(doc models.User) {
		log.Info("User created", doc)
	})
}