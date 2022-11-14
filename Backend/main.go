package main

import (
	"dbms/configuration"
	"dbms/router"
	"dbms/routes"
	"fmt"
)

func init() {
	fmt.Println("DATABASE MANAGEMENT SOFTWARE RUNNING")
}

func main() {
	application := router.NewRouter()
	applicationContext := configuration.Context{}
	routes.HealthCheckRoutes()

	application.Serve(":4000", &applicationContext)
}
