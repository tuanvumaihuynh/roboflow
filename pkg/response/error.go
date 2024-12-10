package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/tuanvumaihuynh/roboflow/pkg/xerrors"
)

// ErrorToResponse converts an error to an ErrorResponse.
func ErrorToResponse(err error) ErrorResponse {
	if err == nil {
		return ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    "Internal server error",
		}
	}

	switch err := err.(type) {
	case *xerrors.AlreadyExistsError:
		return ErrorResponse{
			StatusCode: http.StatusConflict,
			Code:       err.GetCode(),
			Message:    err.GetMessage(),
		}
	case *xerrors.InternalError:
		return ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_server_error",
			Message:    "Internal server error",
		}
	case *xerrors.InvalidArgumentError:
		return ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Code:       err.GetCode(),
			Message:    err.GetMessage(),
		}
	case *xerrors.NotFoundError:
		return ErrorResponse{
			StatusCode: http.StatusNotFound,
			Code:       err.GetCode(),
			Message:    err.GetMessage(),
		}
	case *xerrors.PermissionDeniedError:
		return ErrorResponse{
			StatusCode: http.StatusForbidden,
			Code:       err.GetCode(),
			Message:    err.GetMessage(),
		}
	case *xerrors.UnauthenticatedError:
		return ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Code:       err.GetCode(),
			Message:    err.GetMessage(),
		}
	case *xerrors.UnavailableError:
		return ErrorResponse{
			StatusCode: http.StatusServiceUnavailable,
			Code:       err.GetCode(),
			Message:    err.GetMessage(),
		}
	case validator.ValidationErrors:
		validationErrs := make([]ValidationError, len(err))
		for i, e := range err {
			validationErrs[i] = ValidationError{
				Field:   e.Field(),
				Message: fmt.Sprintf("'%s' %s", e.Field(), e.Tag()),
			}
		}
		return ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Code:       "validation_error",
			Message:    "Validation error",
			Details:    validationErrs,
		}
	default:
		e := new(xerrors.XError)
		if errors.As(err, &e) {
			return ErrorToResponse(errors.Unwrap(err))
		}
		return ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    "Internal server error",
		}
	}
}