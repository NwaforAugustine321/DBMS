package healthcheck

import "net/http"


type heartBeat struct{

}

func NewHealthCheck () HealthCheckInterface{
  return &heartBeat{}
}

func (heart *heartBeat) CheckServer(response http.ResponseWriter,request *http.Request){
	
	response.Write([]byte("Server is alive"))
}