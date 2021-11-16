package helpers

import (
	"example/web-service-gin/dtos"
	"example/web-service-gin/langs"

	"github.com/go-playground/validator/v10"
)

func GenerateValidationResponse(err error) (response dtos.ValidationResponse) {
	response.Success = false

	var valids []dtos.Validation

	validErr := err.(validator.ValidationErrors)

	for _, value := range validErr {
		field, rule := value.Field, value.Tag

		valid := dtos.Validation{Field: field(), Message: langs.GenerateValidationMessage(field(), rule())}

		valids = append(valids, valid)
	}

	response.Validations = valids

	return response
}
