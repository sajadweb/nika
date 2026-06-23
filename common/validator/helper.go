package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/sajadweb/nika/common/response"
)

// ValidateStruct validates a struct using the injected Validator's underlying
// validate instance. Falls back to the global V for backward compatibility.
func ValidateStruct(v *Validator, s interface{}) []FieldError {
	err := v.V.Struct(s)
	if err == nil {
		return nil
	}

	return FormatErrors(err)
}

// BindAndValidate binds JSON body and validates it.
// On failure, responds with an appropriate error and returns false.
func BindAndValidate(v *Validator, c *gin.Context, dto interface{}) bool {
	if err := c.ShouldBindJSON(dto); err != nil {
		response.BadRequest(c, "INVALID_JSON", err.Error())
		return false
	}

	if errs := ValidateStruct(v, dto); errs != nil {
		response.UnprocessableEntity(c, "VALIDATION_ERROR", errs)
		return false
	}

	return true
}

// BindAndValidateQuery binds query parameters and validates them.
func BindAndValidateQuery(v *Validator, c *gin.Context, dto interface{}) bool {
	if err := c.ShouldBindQuery(dto); err != nil {
		response.BadRequest(c, "INVALID_QUERY", err.Error())
		return false
	}

	if errs := ValidateStruct(v, dto); errs != nil {
		response.UnprocessableEntity(c, "VALIDATION_ERROR", errs)
		return false
	}

	return true
}

// BindAndValidateUri binds URI path parameters and validates them.
func BindAndValidateUri(v *Validator, c *gin.Context, dto interface{}) bool {
	if err := c.ShouldBindUri(dto); err != nil {
		response.BadRequest(c, "INVALID_URI", err.Error())
		return false
	}

	if errs := ValidateStruct(v, dto); errs != nil {
		response.UnprocessableEntity(c, "VALIDATION_ERROR", errs)
		return false
	}

	return true
}
