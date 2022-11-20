package userHandlers

import (
	"context"
	"dbms/configuration/databaseConfig"
	"dbms/errorsHandlers"
	"dbms/repo/interfaces"
	"dbms/repo/sql/userRepo"
	"dbms/services"
	"dbms/use-cases/users"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	userRepository = userRepo.NewUserSource(databaseConfig.NewDatabase())
	validation     = services.NewValidation()
	user           = users.NewUser(userRepository, validation)
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var data interfaces.User

	err := json.NewDecoder(request.Body).Decode(&data)

	if err != nil {

		jsonError := &errorsHandlers.JsonErrorParser{
			Message: "Invalid payload",
			Status:  400,
			Err:     err.Error(),
		}

		jsonError.WriteJsonToStream(response)
		logMessage := jsonError.Error()
		fmt.Println(logMessage)
		return
	}

	ctx, cancel := context.WithTimeout(request.Context(), 30*time.Millisecond)
	defer cancel()
	validationError := user.ValidateUserCreation(data)

	if err, ok := validationError.(*errorsHandlers.UserParser); ok {
		if err != nil {
			err.WriteJsonToStream(response)
			return
		}

	}
	
	user.CreateUser(ctx, data)
	response.Write([]byte(" creating user "))
}
