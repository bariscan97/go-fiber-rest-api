package userRoute

import (
	"go-pgx/interval/controller/usersController"
    "github.com/gofiber/fiber/v2"
)


func UserRoute(router fiber.Router) {
	controller := userController.NewUserController()
	user := router.Group("/user")
	user.Get("/me", controller.GetMe)
	user.Delete("/", controller.DeleteMe)
	user.Patch("/", controller.UpdateMe)
}