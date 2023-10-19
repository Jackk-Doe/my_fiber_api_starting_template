package validators

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validatorObj *validator.Validate

func Init() {
	validatorObj = validator.New()
}

func ValidateReqBody(data interface{}) error {
	errs := validatorObj.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			if err != nil {
				// errMsg := fmt.Sprintf("Field '%s' : '%v' | Needs to pass '%s' validation", err.Field(), err.Value(), err.Tag())
				errMsg := fmt.Sprintf("Field '%s' | Needs to pass '%s' validation", err.Field(), err.Tag())
				return errors.New(errMsg)
			}
		}
	}

	return nil
}
