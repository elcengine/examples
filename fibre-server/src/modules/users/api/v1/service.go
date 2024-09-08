package v1

import (
	"fmt"
	"mailman/src/modules/users/api/v1/dto"
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"

	filter_query "github.com/elcengine/elemental/plugins/filter-query"
	filter_query_middleware "github.com/elcengine/elemental/plugins/filter-query/middleware"
	e_utils "github.com/elcengine/elemental/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/sethvargo/go-password/password"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createUser(c *fiber.Ctx, payload dto.CreateUserReq) dto.CreateUserRes {
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
	return dto.CreateUserRes{
		ID:       insertedUser.ID,
		Password: generatedPassword,
	}
}

func getAllUsers(c *fiber.Ctx) dto.GetAllUsersRes {
	users := models.UserModel.Find(e_utils.Cast[filter_query.FilterQueryResult](c.Locals(filter_query_middleware.CTXKey)).Filters).Exec().([]models.User)
	return users
}

func getUserByID(c *fiber.Ctx, id string) dto.GetUserByIDRes {
	user := models.UserModel.FindByID(lo.Must(primitive.ObjectIDFromHex(id))).Exec().(models.User)
	return user
}

func updateUser(c *fiber.Ctx, id string, payload dto.CreateUserReq) dto.UpdateUserRes {
	user := models.UserModel.UpdateByID(lo.Must(primitive.ObjectIDFromHex(id)), models.User{
		Name:          payload.Name,
		Email:         payload.Email,
		Role:          models.UserRole(payload.Role),
		Organizations: payload.Organizations,
	}).Exec().(models.User)
	return user
}

func deleteUser(c *fiber.Ctx, id string) dto.DeleteUserRes {
	user := models.UserModel.FindByIdAndDelete(lo.Must(primitive.ObjectIDFromHex(id))).Exec().(models.User)
	return user
}
