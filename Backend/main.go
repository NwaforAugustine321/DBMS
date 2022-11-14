package main

import (
	"dbms/configuration"
	healthcheck "dbms/healthCheck"
	"dbms/router"
	"fmt"
)

func init(){
	fmt.Println("DATABASE MANAGEMENT SOFTWARE RUNNING")
}

func main(){
   application := router.NewRouter()
   healthCheck := healthcheck.NewHealthCheck()
   applicationContext := configuration.Context{}

   application.Get("/heartbeat",healthCheck.CheckServer)

   application.Serve(":4000",&applicationContext)
}