package healthcheck

import "net/http"

type  HealthCheckInterface interface{
	CheckServer(response http.ResponseWriter,request *http.Request)
}