package userHandlers

import (
	"dbms/errorsHandlers"
	"dbms/repo/interfaces"
	"dbms/use-cases/users"
	"encoding/json"
	"net/http"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var data interfaces.User
	err := json.NewDecoder(request.Body).Decode(&data)

	if err != nil {

		var jsonError = &errorsHandlers.JsonErrorParser{
			Message:   "Failed to parse body ",
			Status: 400,
			Err: err.Error(),
		}
  
		jsonError.WriteJsonToStream(response)
		jsonError.PrintJsonError()

		return
	}

	users.CreateUser(data)
	response.Write([]byte(" creating user "))
}
