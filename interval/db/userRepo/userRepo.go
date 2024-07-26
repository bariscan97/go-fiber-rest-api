package userRepo

import (
	"context"
	"fmt"
	"errors"
	"go-pgx/pkg/models"
	"reflect"
	"strconv"
	"go-pgx/interval/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type (
	UserRepository struct {
		pool  *pgxpool.Pool
	}
	IUserRepo interface {
		CreateUser(data *models.UserModel) error
		UpdateMe(id uuid.UUID , data *models.UpdateUserModel) error
		DeleteMe(id uuid.UUID) error
		GetUserByEmail(email string) (models.UserModel, error)
	}
)

func NewUserRepo() IUserRepo {
	return &UserRepository{
		pool: db.Pool(),
	}
}

func (userRepo *UserRepository) CreateUser(data *models.UserModel) error {
	ctx := context.Background()
	
	sql := `INSERT INTO users(username,email,password) VALUES($1 ,$2, $3)`
	
	commandTag, err := userRepo.pool.Exec(ctx, sql,data.Username,data.Email,data.Password)

	if err != nil {
		
		log.Error("Failed to create new user", err)
		return err 
	}
	log.Info(fmt.Printf("User created with %v", commandTag))
	return nil 
}

func (userRepo *UserRepository)  GetUserByEmail(email string) (models.UserModel, error) {
	
	ctx := context.Background()
    
	sql := `SELECT id , username, email, password FROM users WHERE email = $1`
    
	Rows := userRepo.pool.QueryRow(ctx, sql, email)
    
	var (
		Id 		  uuid.UUID
		Username  string
		Email     string
		Password  string
	)
    
	err := Rows.Scan(&Id, &Username,&Email ,&Password)
    
	if err != nil {
        return models.UserModel{}, err
    }
       
	return models.UserModel{
		Id : Id,
		Username: Username,
		Email: Email,
		Password: Password,
	} , nil
}

func (userRepo *UserRepository)  UpdateMe(id uuid.UUID , data *models.UpdateUserModel) error {
	
	ctx := context.Background()	
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

	var (
        keys string
        sql string
        values []any
		count int
	)
	
	sql = `UPDATE users SET `
	size := v.NumField()

	for i := 0; i< size ; i++ {
        if v.Field(i).Kind() == reflect.String && v.Field(i).String() == "" {
			count++
			continue
		}
		field := v.Type().Field(i).Name + "=" + "$" + strconv.Itoa(i + 1 - count) + ","
        keys += field
        if i == size - 1 {
            keys = keys[:len(keys) - 1] + ` WHERE id = $` + strconv.Itoa(i + 2 - count) 
        }
		
        values = append(values, v.Field(i))
    }
    
    values = append(values, id)
    
	sql += keys
	
	_, err := userRepo.pool.Exec(ctx , sql ,values...)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (userRepo *UserRepository)  DeleteMe(id uuid.UUID) error {
	ctx := context.Background()

	sql := `
		DELETE FROM users
		WHERE id = $1
	`
	commandTag, err := userRepo.pool.Exec(ctx , sql, id)

	
	if err != nil {
		return errors.New(err.Error())
	}
	
	log.Info(fmt.Printf("User created with %v", commandTag))

	return nil
}

