package users

import (
	"context"
	"dbms/repo/interfaces"
	"dbms/repo/sql/userRepo"
	"dbms/services"
)


var userRepository = userRepo.NewUserSource()
var validation = services.NewValidation()

type newUser struct{}

func NewUser( userRepo  interfaces.UserSourceInterface,validation interfaces.UserValidationInterface)interfaces.UserServiceInterface{
  return &newUser{}
}

func (t newUser) CreateUser(ctx context.Context, data interfaces.User){
    userRepository.CreateUser(ctx,data)
}

func (t newUser) ValidateUserCreation( data interfaces.User) error {
    return validation.ValidateUserCreation(data)
}
