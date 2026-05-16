package validator

import (
	"time"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
    validate.RegisterValidation("month_year", validateMonthYear)
}

func validateMonthYear(fl validator.FieldLevel) bool {
    _, err := time.Parse("01-2006", fl.Field().String())
    return err == nil
}

func Validate(s any) error {
    return validate.Struct(s)
}