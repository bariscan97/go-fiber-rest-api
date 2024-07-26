package routes

import (
	"go-pgx/interval/routes/authRoute"
	"go-pgx/interval/routes/todoRoute"
	"go-pgx/interval/routes/userRoute"

	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	authRoute.AuthRoute(app)
	userRoute.UserRoute(app)
	todoRoute.TodoRoute(app)
}
