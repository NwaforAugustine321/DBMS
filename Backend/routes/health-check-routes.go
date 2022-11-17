package routes

import (
	healthcheck "dbms/healthCheck"
	"dbms/router"
)

func HealthCheckRoutes(router router.RouterInterface) {

	router.Get("/health-check/beat", healthcheck.CheckServer)

}
