package authController

import (
	"go-pgx/pkg/models"
	"go-pgx/interval/service/user"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	IAuthController interface {
		Register(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
	}
	AuthController struct {
		UserServices  user_service.IUserService
	}
	UserClaim struct {
		jwt.RegisteredClaims
		Id    		uuid.UUID 
		Email		string
		UserName	string
	}
)

func NewAuthController() IAuthController {
	return &AuthController{
		UserServices : user_service.NewAuthService(),	
	}
}

func (controller *AuthController) Register(c *fiber.Ctx) error  {
	data := new(models.UserModel)
	 
	var err error
	
	if err = c.BodyParser(&data) ; err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	if errs ,notvalid := data.IsValidUser() ; !notvalid {
		return c.JSON(errs)
	}
	bytePassword := []byte(data.Password)
    hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
    if err != nil {
        return fiber.NewError(401,err.Error())
    }
	data.Password = string(hash)
	err = controller.UserServices.CreateUser(data)
	if err != nil {
		return fiber.NewError(404,err.Error())
	}
	return c.JSON(fiber.Map{
		"message" : "successful",
	})
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	
	data := new(models.LoginReqBody)
	
	var err error
	
	if err = c.BodyParser(&data) ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})	
	}
	
	rowResult ,err := controller.UserServices.GetUserByEmail(data.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(rowResult.Password), []byte(data.Password))
	
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	SecretKey := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		Id: rowResult.Id,
		Email: rowResult.Email,
		UserName: rowResult.Username,
	})
	
	tokenString, err := token.SignedString([]byte(SecretKey))
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}
	
	c.Cookie(&fiber.Cookie{
		Name:     "acces_token_jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(72000 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	
	return c.JSON(fiber.Map{
		"err":"Login successful" ,
	})
}
