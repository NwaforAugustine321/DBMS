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
	CreateUser(ctx context.Context, data User)
	
}

type UserServiceInterface interface{
	CreateUser(ctx context.Context, data User)
	ValidateUserCreation(data User)interface{} 
}