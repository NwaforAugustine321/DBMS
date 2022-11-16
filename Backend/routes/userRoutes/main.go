package userRoutes

import (
	"dbms/handlers/userHandlers"
	"dbms/router"
)


func NewUserRoutes(router router.RouterInterface){
   pathPrefix := 	router.PathPrefix("/user")

   pathPrefix.HandleFunc("/create",userHandlers.CreateUser).Methods("POST")
}