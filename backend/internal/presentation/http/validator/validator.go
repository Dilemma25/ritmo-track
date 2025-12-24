package validator

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"ritmotrack-backend/internal/presentation/http/apierror"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	vWrapper := &Validator{
		validate: v,
	}

	vWrapper.registerValidation("regex", regexValidation)

	return vWrapper
}

func (ths *Validator) registerValidation(tag string, fn validator.Func) {
	if err := ths.validate.RegisterValidation(tag, fn); err != nil {
		log.Fatalf("failed to registration validation: %v", err)
	}
}

func (ths *Validator) ValidateBody(r *http.Request, d any) apierror.ApiError {
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		return apierror.NewErrInvalidJson(err)
	}

	if err := ths.validate.Struct(d); err != nil {

		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			errMap := make(map[string]string)

			for _, fe := range verr {
				field := jsonFieldName(d, fe)
				errMap[field] = validationErrorMessage(fe)
			}

			return apierror.NewErrValidationFailed(errMap)
		}
	}

	return nil
}

// Получаем название JSON-поля по ошибке валидации
func jsonFieldName(d any, fe validator.FieldError) string {
	t := reflect.TypeOf(d).Elem() // Тип структуры
	if f, ok := t.FieldByName(fe.StructField()); ok {
		tag := strings.Split(f.Tag.Get("json"), ",")[0]

		name := strings.Split(tag, ",")[0]

		if name == "" || name == "-" {
			return fe.StructField() // Fallback на имя поля структуры
		}

		return tag
	}

	return fe.StructField()
}

//Customise validation

func regexValidation(fl validator.FieldLevel) bool {
	pattern := fl.Param()
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(fl.Field().String())
}
