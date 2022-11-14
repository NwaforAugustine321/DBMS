package routes

import (
	healthcheck "dbms/healthCheck"
	"dbms/router"
)

func HealthCheckRoutes() {
	router := router.NewRouter()
	healthCheck := healthcheck.NewHealthCheck()

	routerPrefixPath := router.PathPrefix("/health-check")

	routerPrefixPath.HandleFunc("/beat", healthCheck.CheckServer).Methods("GET")
	

}
