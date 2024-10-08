package v1

import (
	"fmt"
	"mailman/src/modules/auth/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func loginWithEmailAndPassword(c *fiber.Ctx, email, password string) *dto.LoginRes {
	user := models.UserModel.FindOne(primitive.M{"email": email}).Exec().(models.User)
	if user.ID == primitive.NilObjectID {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials"))
	}
	passwordsMatch := utils.CompareStrHash(password, user.Password)
	if !passwordsMatch {
		panic(fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials"))
	}
	accessToken := utils.GenerateUserJWTToken(user, false)
	refreshToken := utils.GenerateUserJWTToken(user, true)
	return &dto.LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user.Secure(),
	}
}

func registerUser(c *fiber.Ctx, payload dto.RegisterReq) *dto.LoginRes {
	user := models.UserModel.Create(models.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Age:      payload.Age,
		Verified: true,
		Password: utils.HashStr(payload.Password),
	}).Exec().(models.User)
	log.Info(fmt.Sprintf("Account creation successful. Email - %s Generating tokens....", payload.Email), user.ID)
	accessToken := utils.GenerateUserJWTToken(user, false)
	refreshToken := utils.GenerateUserJWTToken(user, true)
	return &dto.LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}
}
