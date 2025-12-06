package validation

import (
	"encoding/json"
	"errors"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translator, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, translator)
	}
}

func ValidateUserError(validation_err error) *rest_error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_error.NewBadRequestError("Invalid field type")
	}

	if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_error.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_error.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return rest_error.NewBadRequestEValidationError("Some fields are invalid", errorsCauses)
	}

	return rest_error.NewBadRequestError("Error try to convert fields")
}
