package v1

import (
	"fmt"
	"mailman/src/modules/users/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
)

func createUser(c *fiber.Ctx, payload dto.CreateUserReq) *dto.CreateUserRes {
	log.Info("Creating user within system - ", payload.Email)
	verificationCode := uuid.New().String()
	generatedPassword, _ := password.Generate(8, 2, 1, false, false)
	fmt.Println(payload.Name, payload.Email, generatedPassword, verificationCode)
	insertedUser := models.UserModel.Create(models.User{
		Email:            payload.Email,
		Name:             payload.Name,
		VerificationCode: &verificationCode,
		Password:         utils.HashStr(generatedPassword),
	}).Exec().(models.User)
	return &dto.CreateUserRes{
		ID:       insertedUser.ID,
		Password: generatedPassword,
	}
}
