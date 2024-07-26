package todoController

import (
	"go-pgx/interval/service/todo"
	"go-pgx/pkg/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoController struct {
	todoService todo_service.ITodoService
}
type ITodoController interface {
	CreateTodo(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
	GetTodoById(c *fiber.Ctx) error
	GetAllTodos(c *fiber.Ctx) error
}

func NewTodoController() ITodoController {
	return &TodoController{
		todoService: todo_service.NewTodoService(),
	}
}

func (todoController *TodoController) CreateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.MyCustomClaims)
	data := new(models.CreateTodo)
	var (
		err error
	)
	if err = c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	err = todoController.todoService.CreateTodo(user.Id, data.Content)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  err.Error(),
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message" : "successful",
	})
}

func (todoController *TodoController)  UpdateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.MyCustomClaims)
	data := new(models.CreateTodo)
	var (
		err error
	)
	todo_id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.JSON(fiber.Map{
			"error" : err.Error(),
		})
    }
	if err = c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	err = todoController.todoService.UpdateTodo(user.Id,todo_id ,data.Content)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  err.Error(),
		})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message" : "successful",
	})
}
func (todoController *TodoController)  DeleteTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.MyCustomClaims)
	
	todo_id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.JSON(fiber.Map{
			"error" : err.Error(),
		})
    }
	err = todoController.todoService.DeleteTodo(user.Id , todo_id)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message" : "successful",
	})
}
func (todoController *TodoController)  GetTodoById(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.MyCustomClaims)
	var err error
	todo_id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.JSON(fiber.Map{
			"error" : err.Error(),
		})
    }
	rows, err := todoController.todoService.GetTodoById(user.Id, todo_id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  err.Error(),
		})
	} 

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"result" : rows,
	})

}

func (todoController *TodoController)  GetAllTodos(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.MyCustomClaims)
	var err error
	page := c.Query("page")
	
	if page == "" {
		page = "0"
	}
	
	intPage ,err:= strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  err.Error(),
		})
	}
	
	rowsResult ,err := todoController.todoService.GetAllTodos(user.Id, intPage)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"result" : rowsResult,
	})
}