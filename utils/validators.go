package utils

import (
	"yas/types"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) []*types.ErrorResponse {
	var errors []*types.ErrorResponse
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &types.ErrorResponse{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			})
		}
	}
	return errors
}
