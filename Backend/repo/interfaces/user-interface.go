package interfaces

import (
	"context"
)

type User struct {
	UserName  string `json:"userName" validate:"required"`
	UserEmail string `json:"userEmail" validate:"required,email"`
	UserPhone string `json:"userPhone"`
	Profile   string `json:"profile"`
}

type UserSourceInterface interface {
	CreateUser(ctx context.Context, data User)(error,string)
	GetAllUser(ctx context.Context)([]User,error)
}

type UserServiceInterface interface{
	CreateUser(ctx context.Context, data User)(error,string)
	GetAllUser(ctx context.Context)([]User,error)
	ValidateUserCreation(data User)error
}