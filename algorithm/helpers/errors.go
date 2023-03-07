package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/x/multiservice/algorithm/types/requests"
)

// GetErrorMsg returns a string describing the error from a request payload
//
// It takes a validator.FieldError as input and returns a string
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "optional":
		return ""
	case "email":
		return "This field must be a valid email address"
	case "oneof":
		return "This field must be one of password, refresh_token, otp_request, or otp_token if grant_type or one of regularAccount, serviceProvider if profileAccountType"
	case "required":
		return "This field is required field and has to be provided"
	case "alpha":
		return "This field should contain only english alphabets"
	case "lte":
		return "This field should be less than " + fe.Param()
	case "gte":
		return "This field should be greater than " + fe.Param()
	case "min":
		return "This field should not be less than " + fe.Param()
	case "max":
		return "This field should not be greater than " + fe.Param()
	}
	return "Unknown error"
}

// FormatError formats an error message
//
// It takes a context object and a an error string as input and aborts the operation with
// bad request status and the error message in the payload
func FormatError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]requests.ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = requests.ErrorMsg{
				Field:   fe.Field(),
				Message: GetErrorMsg(fe),
			}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	}
}
