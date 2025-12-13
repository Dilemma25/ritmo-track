package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	v := validator.New()

	err := v.RegisterValidation("regex", regexValidation)

	if err != nil {
		panic(err)
	}

	return &Validator{
		validate: v,
	}
}

// TODO сделать адекватный вывод ошибок валидации
func (ths *Validator) ValidateStruct(data interface{}) map[string]string {
	errs := map[string]string{}

	if err := ths.validate.Struct(data); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Field()] = fmt.Sprintf("failed on '%s' tag in '%s'", e.Tag(), e.Field())
		}
	}

	return errs
}

//Customise validation

func regexValidation(fl validator.FieldLevel) bool {
	pattern := fl.Param()
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(fl.Field().String())
}
