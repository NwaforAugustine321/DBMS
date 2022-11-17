package userRoutes

import (
	"dbms/handlers/userHandlers"
	"dbms/router"
)


func NewUserRoutes(router router.RouterInterface){
   
   router.Post("/user/create",userHandlers.CreateUser)
  
}