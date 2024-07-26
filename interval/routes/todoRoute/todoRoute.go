package todoRoute


import (
	"github.com/gofiber/fiber/v2"
	"go-pgx/interval/controller/todoController"
)


func TodoRoute(router fiber.Router) {
	controller := todoController.NewTodoController()
	todo := router.Group("/todo")
	todo.Get("/" ,controller.GetAllTodos)
	todo.Post("/", controller.CreateTodo)
	todo.Get("/:id" ,controller.GetTodoById)
	todo.Delete("/:id" ,controller.DeleteTodo)
	todo.Patch("/:id", controller.UpdateTodo)
	
}


