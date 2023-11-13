package validation

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type validationErrorResponse struct {
	Tag   string
	Value interface{}
}

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}

func (v Validator) Decode(errs error) *map[string]interface{} {

	validationErrors := map[string]validationErrorResponse{}
	for _, err := range errs.(validator.ValidationErrors) {
		validationErrors[err.Field()] = validationErrorResponse{
			Tag:   err.Tag(),
			Value: err.Value(),
		}
	}
	return &map[string]interface{}{"errors": &validationErrors}
}

func (v Validator) JSON(errs error) *[]byte {
	errsMap := v.Decode(errs)

	res, _ := json.Marshal(errsMap)

	return &res
}
