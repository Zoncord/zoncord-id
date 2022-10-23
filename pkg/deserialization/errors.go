package deserialization

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Field  string `json:"field"`
	Detail string `json:"detail"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "oneof":
		return "The field must have one of the following values: " + fe.Param()
	case "required_if":
		// TODO give more context
		return "This field is required"
	}
	return "Unknown Error"
}

func GetDetailedErrors(err error) []Error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]Error, len(ve))
		for i, fe := range ve {
			out[i] = Error{Field: fe.Field(), Detail: getErrorMsg(fe)}
		}
		return out
	}
	return []Error{}
}
