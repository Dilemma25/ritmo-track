package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validationFuncs = map[string]func(fe validator.FieldError) string{
	"required": func(fe validator.FieldError) string { return "Обязательное поле" },
	"email":    func(fe validator.FieldError) string { return "Неверный формат email" },
	"min": func(fe validator.FieldError) string {
		return fmt.Sprintf("Минимальная длинна - %s", fe.Param())
	},
	"max": func(fe validator.FieldError) string {
		return fmt.Sprintf("Максимальная длинна - %s", fe.Param())
	},
	"lte": func(fe validator.FieldError) string {
		return fmt.Sprintf("Максимальное значение - %s", fe.Param())
	},
	"gte": func(fe validator.FieldError) string {
		return fmt.Sprintf("Минимальное значение - %s", fe.Param())
	},
	"lt": func(fe validator.FieldError) string {
		return fmt.Sprintf("Значение должно быть меньше %s", fe.Param())
	},
	"gt": func(fe validator.FieldError) string {
		return fmt.Sprintf("Значение должно быть больше %s", fe.Param())
	},
}

func validationErrorMessage(fe validator.FieldError) string {
	if f, ok := validationFuncs[fe.Tag()]; ok {
		return f(fe)
	}

	return "Неверное значение"
}
