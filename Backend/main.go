package main

import (
	"dbms/configuration"
	"dbms/configuration/databaseConfig"
	"dbms/errorsHandlers"
	"dbms/router"
	"dbms/routes"
	"dbms/routes/invoiceRoutes"
	"dbms/routes/userRoutes"
	"fmt"
	"log"
)

func init() {
	fmt.Println("INVOICE GENERATOR SOFTWARE RUNNING")
}

const (
	dsn = "root:@Root321@tcp(127.0.0.1:3306)/AtlasDB?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {

	// SET UP APPLICATION AND CONTEXT CONFIGURATIONS
	applicationRouter := router.NewRouter()
	applicationContext := configuration.Context{}

	//SET UP APIS ROUTES
	routes.HealthCheckRoutes(applicationRouter)
	userRoutes.NewUserRoutes(applicationRouter)
	invoiceRoutes.NewInvoiceRoutes(applicationRouter)

	//INIT DATABASE CONNECTION
	db := databaseConfig.NewDatabase()
	err := db.InitDB(dsn)
	if err, ok := err.(*errorsHandlers.DataBaseError); ok {
		if err != nil {
			log.Fatal(err.Error())
		}

	}

	err = db.Migrate()

	if err, ok := err.(*errorsHandlers.DataBaseError); ok {
		if err != nil {
			log.Fatal(err.Error())
		}

	}

	
	applicationRouter.Serve(":4000", &applicationContext)
}
