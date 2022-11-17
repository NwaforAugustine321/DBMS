package healthcheck

import "net/http"



func  CheckServer(response http.ResponseWriter,request *http.Request){
	
	response.Write([]byte("Server is alive"))
}