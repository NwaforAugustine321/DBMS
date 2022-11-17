package validation

import (
	"dbms/repo/interfaces"

	"github.com/go-playground/validator/v10"
)

type validation struct{}

var (
	validate = validator.New()
)

func NewValidation() interfaces.ValidationInterface {
	return &validation{}
}

func (t *validation) ValidateStruct(data interface{}) {

	
}

func (t *validation) ValidateKey() {}
