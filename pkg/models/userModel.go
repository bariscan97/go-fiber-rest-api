package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Valid interface {
	IsValidUser() (map[string]string, bool)
}

type UserModel struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username" validate:"required,min=5,max=20"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=24"`
}

type LoginReqBody struct {
	Email    string
	Password string
}

type UpdateUserModel struct {
	Username string `json:"username" validate:"omitempty,min=5,max=20"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=8,max=24"`
}

func validateStruct(v interface{}) (map[string]string, bool) {
	validate := validator.New()
	err := validate.Struct(v)
	errors := make(map[string]string)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}
		return errors, false
	}
	return errors, true
}

func (u *UserModel) IsValidUser() (map[string]string, bool) {
	return validateStruct(u)
}

func (u *UpdateUserModel) IsValidUser() (map[string]string, bool) {
	return validateStruct(u)
}

func ValidCheck(u Valid) {
	u.IsValidUser()
}
