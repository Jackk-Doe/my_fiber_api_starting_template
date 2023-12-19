package validators

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validatorObj *validator.Validate

func Init() {
	validatorObj = validator.New()
}

/*
Validate struct with tag, of request body, param, query, etc.
*/
func validateStructWithTags(data interface{}) error {
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

/*
Validate uuid string input
*/
func ValidateUuid(uuid string) error {
	err := validatorObj.Var(uuid, "required,uuid")
	if err != nil {
		return errors.New("Invalid uuid")
	}

	return nil
}

/*
Parse and validate body
*/
func ParseAndValidateBody(c *fiber.Ctx, out interface{}) error {
	if err := c.BodyParser(out); err != nil {
		return errors.New(err.Error())
	}
	return validateStructWithTags(out)
}

/*
Parse and validate params
*/
func ParseAndValidateQueryParam(c *fiber.Ctx, out interface{}) error {
	if err := c.QueryParser(out); err != nil {
		return errors.New(err.Error())
	}
	return validateStructWithTags(out)
}
