package validate

import (
	"fmt"

	"github.com/go-playground/validator"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(T interface{}) error {
	err := validate.Struct(T)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%s %v %v %s", "field", err.Field(), err.Value(), "is illegal")
		}
	}

	return nil
}
