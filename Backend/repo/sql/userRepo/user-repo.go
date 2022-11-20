package userRepo

import (
	"context"
	"dbms/models"
	"dbms/repo/interfaces"
	"fmt"
)

type userSource struct{
	db interfaces.DatabaseInterface
}

func NewUserSource(db interfaces.DatabaseInterface) interfaces.UserSourceInterface {
	return &userSource{db}
}

func (user *userSource) CreateUser(ctx context.Context, data interfaces.User) {
	// var userModel models.Users
	// userModel.UserName = data.UserName
	// userModel.UserEmail = data.UserEmail
	// userModel.UserPhone = data.UserPhone
	// userModel.Profile = data.Profile
	// userModel.UserID = uuid.UUID{}
	// databaseConfig.NewDatabase().DBInstance().Create(&userModel)
	var result  models.Users
   var r = user.db.Get("Users",[]string{"user_name","user_email"},"WHERE user_name = ? AND user_email = ?",&result, "austine","nea@gmail.com")
   fmt.Println(r)
}

	

