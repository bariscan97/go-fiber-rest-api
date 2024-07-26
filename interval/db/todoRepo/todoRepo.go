package todoRepo

import (
	"context"
	"go-pgx/interval/db"
	"go-pgx/pkg/models"
	"time"
    "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TodoRepository struct {
	pool *pgxpool.Pool
}
type ITodoRepository interface {
	CreateTodo(user_id uuid.UUID, content string) error 
	UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) error
	DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) error 
	GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error)
	GetAllTodos(user_id uuid.UUID ,page int) ([]models.FetchTodoModel, error)
}

func NewUserRepo() ITodoRepository {
	return &TodoRepository{
		pool: db.Pool(),
	}
}

func (todoRepo *TodoRepository) CreateTodo(user_id uuid.UUID, content string) error {
	ctx := context.Background()

	sql := `INSERT INTO todos(user_id ,content) VALUES($1 ,$2)`

	_, err := todoRepo.pool.Exec(ctx, sql, user_id, content)

	if err != nil {
		return err
	}

	return nil

}

func (todoRepo *TodoRepository) UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) error {
	ctx := context.Background()
	
	sql := `
		UPDATE todos
		SET 
			content = $1
		WHERE 
			user_id = $2 AND id = $3

		`
	_, err := todoRepo.pool.Exec(ctx, sql, content, user_id, todo_id)
	
	if err != nil {
		return err
	}

	return nil
}

func (todoRepo *TodoRepository) DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) error {
	ctx := context.Background()

	sql := `
		DELETE FROM todos
		WHERE user_id = $1 AND id = $2
		`
	_, err := todoRepo.pool.Exec(ctx, sql, user_id, todo_id)

	if err != nil {
		return err
	}

	return nil
}

func (todoRepo *TodoRepository) GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error) {
	ctx := context.Background()

	sql := `
		SELECT * FROM todos 
		WHERE user_id = $1 AND id = $2
	`
	rows := todoRepo.pool.QueryRow(ctx, sql, user_id, todo_id)

	var (
		id       uuid.UUID
		userID   uuid.UUID 
		content  string
		createAt time.Time
	)

	err := rows.Scan(&id, &content, &userID, &createAt)

	if err != nil {
		return models.FetchTodoModel{}, err
	}

	return models.FetchTodoModel{
		Id: id,
		Content: content,
		CreateAt : createAt,
	}, nil
}

func (todoRepo *TodoRepository) GetAllTodos(user_id uuid.UUID ,page int) ([]models.FetchTodoModel, error) {
	ctx := context.Background()

	sql := `
		SELECT id, content ,user_id ,created_at FROM todos 
		WHERE user_id = $1
		
	`
	rows ,err := todoRepo.pool.Query(ctx, sql, user_id)
	if err != nil {
		return []models.FetchTodoModel{} , err
	}
	
	var Todos []models.FetchTodoModel

	for rows.Next() {
		var (
			id       uuid.UUID
			content  string
			userID   uuid.UUID 
			createAt time.Time
			DbError  error
		)
		DbError = rows.Scan(&id, &content, &userID, &createAt)
		if DbError != nil {
			return []models.FetchTodoModel{} , DbError
		}
		Todos = append(Todos, models.FetchTodoModel{
			Id: id,
			Content: content,
			CreateAt: createAt,
		})
	}
	return Todos  ,nil
}



// CREATE TABLE todos (
//     id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
//     content VARCHAR(255) NOT NULL,
//     user_id UUID NOT NULL,
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// 	   FOREIGN KEY(user_id) REFERENCES users(id)
// );