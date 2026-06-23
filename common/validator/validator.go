package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/sajadweb/nika"
)

type Validator struct {
	V *validator.Validate
}

type Config struct {
	// Reserved for future validator-level config
}

// Setup creates a new Validator instance, registers custom validations,
// and registers it in the DI container.
func Setup(app *nika.App, cfg Config) *Validator {
	v := validator.New()

	_ = v.RegisterValidation("ir_mobile", validateIRMobile)
	_ = v.RegisterValidation("objectid", validateObjectid)

	validator := &Validator{V: v}

	app.RegisterSingleton(validator)

	fmt.Println("✅ Validator initialized")
	return validator
}

// Set registers an additional custom validation tag.
func (v *Validator) Set(tag string, fn validator.Func) error {
	return v.V.RegisterValidation(tag, fn)
}

func validateIRMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	match, _ := regexp.MatchString(`^09\d{9}$`, mobile)
	return match
}

func validateObjectid(fl validator.FieldLevel) bool {
	objectid := fl.Field().String()
	match, _ := regexp.MatchString(`^[a-f0-9]{24}$`, objectid)
	return match
}
