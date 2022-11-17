package userHandlers

import (
	"dbms/errorsHandlers"
	"dbms/repo/interfaces"
	"dbms/services"
	"dbms/use-cases/users"
	"encoding/json"
	"net/http"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var data interfaces.User

	err := json.NewDecoder(request.Body).Decode(&data)

	if err != nil {

		var jsonError = &errorsHandlers.JsonErrorParser{
			Message: "Invalid payload",
			Status:  400,
			Err:     err.Error(),
		}

		jsonError.WriteJsonToStream(response)
		jsonError.PrintJsonError()
		return
	}

	services.ValidateStruct(data)
	users.CreateUser(data)
	response.Write([]byte(" creating user "))
}
