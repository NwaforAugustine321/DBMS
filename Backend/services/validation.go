package services

import "dbms/repo/validation"


var validate = validation.NewValidation()


func ValidateStruct(data interface{}) {
	 validate.ValidateStruct(data)
}