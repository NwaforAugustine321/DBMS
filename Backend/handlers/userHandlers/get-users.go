package userHandlers

import (
	"context"
	"dbms/configuration/databaseConfig"
	"dbms/errorsHandlers"
	"dbms/repo/sql/userRepo"
	"dbms/services"
	"dbms/use-cases/users"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	getUsers = users.NewUser(userRepo.NewUserSource(databaseConfig.NewDatabase()), services.NewValidation())
)

func GetAllUser(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), 30*time.Millisecond)

	defer cancel()

	data, err := getUsers.GetAllUser(ctx)

	if err != nil {
		return
	}

	responseData, err := json.Marshal(data)

	if err != nil {
		jsonError := &errorsHandlers.JsonErrorParser{
			Message: "Internal server error",
			Status:  500,
			Err:     err.Error(),
		}

		jsonError.WriteJsonToStream(response)
		logMessage := jsonError.Error()
		fmt.Println(logMessage)
		return
	}

	_, responseErr := response.Write(responseData)

	if  responseErr != nil {
		jsonError := &errorsHandlers.JsonErrorParser{
			Message: "Internal server error",
			Status:  500,
			Err:      responseErr.Error(),
		}

		jsonError.WriteJsonToStream(response)
		logMessage := jsonError.Error()
		fmt.Println(logMessage)
		return
	}
}
