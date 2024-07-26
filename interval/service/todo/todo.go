package todo_service


import (
	"go-pgx/interval/db/todoRepo"
	"github.com/google/uuid"
	"go-pgx/pkg/models"
)

type TodoService struct {
	todoRepo todoRepo.ITodoRepository
}
type ITodoService interface {
	CreateTodo(user_id uuid.UUID, content string) error 
	UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) error
	DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) error 
	GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error)	
	GetAllTodos(user_id uuid.UUID, page int) ([]models.FetchTodoModel, error)		
}

func NewTodoService() ITodoService {
	return &TodoService {
		todoRepo: todoRepo.NewUserRepo(),	
	}
}

func (todoService *TodoService) CreateTodo(user_id uuid.UUID, content string) error {
	err := todoService.todoRepo.CreateTodo(user_id, content)

	if err != nil {
		return err
	}
	
	return nil
}

func (todoService *TodoService) UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) error {
	err := todoService.todoRepo.UpdateTodo(user_id, todo_id, content)

	if err != nil {
		return err
	}

	return nil
}
func (todoService *TodoService) DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) error {
	err := todoService.todoRepo.DeleteTodo(user_id, todo_id)

	if err != nil {
		return err
	}

	return nil
}
func (todoService *TodoService) GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error) {
	result ,err := todoService.todoRepo.GetTodoById(user_id ,todo_id)
	
	if err != nil {
		return  models.FetchTodoModel{}, err
	}
	
	return result ,nil
}

func (todoService *TodoService) GetAllTodos(user_id uuid.UUID, page int) ([]models.FetchTodoModel, error) {
	result , err := todoService.todoRepo.GetAllTodos(user_id, page)
	if err != nil {
		return []models.FetchTodoModel{}, err
	}
	return result ,nil
}