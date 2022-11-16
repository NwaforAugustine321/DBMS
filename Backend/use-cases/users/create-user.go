package users

import (
	"dbms/repo/interfaces"
	"dbms/repo/sql/userRepo"
)


var userRepository = userRepo.NewUserSource()

func CreateUser(userDetails interfaces.User){
  userRepository.CreateUser(userDetails)
}