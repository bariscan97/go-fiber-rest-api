package user_service

import (
	"go-pgx/interval/db/userRepo"
	"go-pgx/pkg/models"

	"github.com/google/uuid"
)

type (
	IUserService interface {
		CreateUser(data *models.UserModel) error
		GetUserByEmail(email string) (models.UserModel, error)
		UpdateMe(id uuid.UUID, data *models.UpdateUserModel) error
		DeleteMe(id uuid.UUID) error
	}
	UserService struct {
		userRepo userRepo.IUserRepo
	}
)

func NewAuthService() IUserService {
	return &UserService{
		userRepo: userRepo.NewUserRepo(),
	}
}

func (userService *UserService) CreateUser(data *models.UserModel) error {
	if err := userService.userRepo.CreateUser(data); err != nil {
		return err
	}
	return nil
}

func (userService *UserService) GetUserByEmail(email string) (models.UserModel, error) {
	result, err := userService.userRepo.GetUserByEmail(email)

	if err != nil {
		return models.UserModel{}, err
	}
	return result, nil
}

func (userService *UserService) UpdateMe(id uuid.UUID, data *models.UpdateUserModel) error {
	err := userService.userRepo.UpdateMe(id, data)

	if err != nil {
		return nil
	}
	return nil
}
func (userService *UserService) DeleteMe(id uuid.UUID) error {
	err := userService.userRepo.DeleteMe(id)

	if err != nil {
		return err
	}

	return nil
}
