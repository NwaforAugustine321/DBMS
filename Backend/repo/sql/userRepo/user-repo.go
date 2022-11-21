package userRepo

import (
	"context"
	"dbms/errorsHandlers"
	"dbms/models"
	"dbms/repo/interfaces"

	"github.com/google/uuid"
)

type userSource struct {
	DBRepository interfaces.DatabaseInterface
}

func NewUserSource(db interfaces.DatabaseInterface) interfaces.UserSourceInterface {
	return &userSource{db}
}

func (user *userSource) CreateUser(ctx context.Context, data interfaces.User) (error, string) {
	var existingUser []models.Users
	var userModel models.Users

	userModel.UserName = data.UserName
	userModel.UserEmail = data.UserEmail
	userModel.UserPhone = data.UserPhone
	userModel.Profile = data.Profile
	userModel.UserID = uuid.UUID{}

	user.DBRepository.DBInstance().Where("user_email = ?", data.UserEmail).First(&existingUser)

	if length := len(existingUser); length > 0 {

		return &errorsHandlers.DataBaseError{
			Status:  400,
			Message: "User already exit",
		}, ""
	}

	user.DBRepository.DBInstance().Create(&userModel)
	return nil, "User created successfully"

}



func (user *userSource) GetAllUser(ctx context.Context) ([]interfaces.User , error) {
	var userModel []interfaces.User

	user.DBRepository.DBInstance().Find(&userModel)

	return userModel, nil
}
