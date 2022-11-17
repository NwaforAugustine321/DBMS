package userHandlers

import (
	"context"
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

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var data interfaces.User
	var userRepository = userRepo.NewUserSource()
	var validation = services.NewValidation()
	user := users.NewUser(userRepository, validation)

	err := json.NewDecoder(request.Body).Decode(&data)

	if err != nil {

		var jsonError = &errorsHandlers.JsonErrorParser{
			Message: "Invalid payload",
			Status:  400,
			Err:     err.Error(),
		}

		jsonError.WriteJsonToStream(response)
		jsonError.Error()
		return
	}

	ctx, cancel := context.WithTimeout(request.Context(), 30 * time.Millisecond)
	defer cancel()
	userValidationError := user.ValidateUserCreation(data)
	fmt.Println(userValidationError)

	// if err, ok := validationError.(*errorsHandlers.UserParser);ok {
	// 	if err != nil {
	// 		err.WriteJsonToStream(response)
	// 		return
	// 	}
       
	// }
	user.CreateUser(ctx, data)
	response.Write([]byte(" creating user "))
}
