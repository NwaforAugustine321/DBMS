package services

import (
	"dbms/errorsHandlers"
	"dbms/repo/interfaces"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type userValidation struct{}
type t struct {}

func NewValidation() interfaces.UserValidationInterface {
	return &userValidation{}
}

func (source *userValidation) ValidateUserCreation(user interfaces.User) interface{} {
	err := validation.Errors{
		"name":        validation.Validate(user.UserName, validation.Required, validation.Length(5, 20), is.Alphanumeric.Error("must be a valid name")),
		"email":       validation.Validate(user.UserEmail, validation.Required, is.Email.Error("must be a valid email address")),
		"phoneNumber": validation.Validate(user.UserPhone, validation.Required, is.Digit.Error("must be a valid phone number")),
	}.Filter()

	
	userValidationError := errorsHandlers.UserParser{
		Status: 400,
        Message: err.Error(),
		
	}

	if err != nil {
		return userValidationError
	}
	
	return &t{}
}
	
