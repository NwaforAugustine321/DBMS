package userRepo

import "dbms/repo/interfaces"


type userSource struct{}

func NewUserSource()interfaces.UserSourceInterface{
   return &userSource{}	
}

func (user *userSource) CreateUser(data interfaces.User){

}