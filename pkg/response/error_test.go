package response_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"github.com/tuanvumaihuynh/roboflow/pkg/response"
	"github.com/tuanvumaihuynh/roboflow/pkg/xerrors"
)

// FieldErrorMock is a mock implementation of validator.FieldError for testing purposes.
type mockFieldError struct {
	validator.FieldError
	tag   string
	field string
}

func (e mockFieldError) Tag() string { return e.tag }

func (e mockFieldError) Field() string { return e.field }

func TestErrorToResponse(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected response.ErrorResponse
	}{
		{
			name: "nil error",
			err:  nil,
			expected: response.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Code:       "internal_error",
				Message:    "Internal server error",
			},
		},
		{
			name: "AlreadyExistsError",
			err:  xerrors.ThrowAlreadyExists(nil, "Already exists"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusConflict,
				Code:       "already_exists",
				Message:    "Already exists",
			},
		},
		{
			name: "InternalError",
			err:  &xerrors.InternalError{},
			expected: response.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Code:       "internal_server_error",
				Message:    "Internal server error",
			},
		},
		{
			name: "InvalidArgumentError",
			err:  xerrors.ThrowInvalidArgument(nil, "Invalid argument"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Code:       "invalid_argument",
				Message:    "Invalid argument",
			},
		},
		{
			name: "NotFoundError",
			err:  xerrors.ThrowNotFound(nil, "Not found"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusNotFound,
				Code:       "not_found",
				Message:    "Not found",
			},
		},
		{
			name: "PermissionDeniedError",
			err:  xerrors.ThrowPermissionDenied(nil, "Permission denied"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusForbidden,
				Code:       "permission_denied",
				Message:    "Permission denied",
			},
		},
		{
			name: "UnauthenticatedError",
			err:  xerrors.ThrowUnauthenticated(nil, "Unauthenticated"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Code:       "unauthenticated",
				Message:    "Unauthenticated",
			},
		},
		{
			name: "UnavailableError",
			err:  xerrors.ThrowUnavailable(nil, "Service unavailable"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusServiceUnavailable,
				Code:       "unavailable",
				Message:    "Service unavailable",
			},
		},
		{
			name: "Unknown error type",
			err:  errors.New("unknown error"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Code:       "internal_error",
				Message:    "Internal server error",
			},
		},
		{
			name: "XError wrapped",
			err:  xerrors.ThrowError(errors.New("wrapped error"), "Outer error"),
			expected: response.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Code:       "internal_error",
				Message:    "Internal server error",
			},
		},
		{
			name: "ValidationError",
			err: validator.ValidationErrors{
				&mockFieldError{
					tag:   "required",
					field: "Field1",
				},
				&mockFieldError{
					tag:   "email",
					field: "Field2",
				},
			},
			expected: response.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Code:       "validation_error",
				Message:    "Validation error",
				Details: []response.ValidationError{
					{Field: "Field1", Message: "'Field1' required"},
					{Field: "Field2", Message: "'Field2' email"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := response.ErrorToResponse(tt.err)
			assert.Equal(t, tt.expected, result)
		})
	}
}