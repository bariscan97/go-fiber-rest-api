package authRoute

import (
	"go-pgx/interval/controller/authController"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(router fiber.Router) {
	controller := authController.NewAuthController()
	auth := router.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
}
