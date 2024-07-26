package userController

import (
	user_service "go-pgx/interval/service/user"
	"go-pgx/pkg/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserServices user_service.IUserService
}
type IUserController interface {
	GetMe(c *fiber.Ctx) error
	DeleteMe(c *fiber.Ctx) error
	UpdateMe(c *fiber.Ctx) error
}

func NewUserController() IUserController {
	return &UserController{
		UserServices: user_service.NewAuthService(),
	}
}

func (controller *UserController) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.MyCustomClaims)
	return c.JSON(user)

}

func (controller *UserController) DeleteMe(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.MyCustomClaims)

	err := controller.UserServices.DeleteMe(user.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"delete": "successful",
	})
}

func (controller *UserController) UpdateMe(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.MyCustomClaims)
	var err error
	data := new(models.UpdateUserModel)

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if errs, notvalid := data.IsValidUser(); !notvalid {
		return c.JSON(errs)
	}

	if data.Password != "" {
		bytePassword := []byte(data.Password)
		hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
		if err != nil {
			return fiber.NewError(401, err.Error())
		}
		data.Password = string(hash)

	}
	err = controller.UserServices.UpdateMe(user.Id, data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"update": "successful",
	})
}
