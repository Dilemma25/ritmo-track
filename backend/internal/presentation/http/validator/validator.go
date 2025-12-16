package validator

import (
	"fmt"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	v := validator.New()

	vWrapper := &Validator{
		validate: v,
	}

	vWrapper.RegisterValidation("regex", regexValidation)

	return vWrapper
}

func (ths *Validator) RegisterValidation(tag string, fn validator.Func) {
	if err := ths.validate.RegisterValidation(tag, fn); err != nil {
		log.Fatalf("failed to registration validation: %v", err)
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
