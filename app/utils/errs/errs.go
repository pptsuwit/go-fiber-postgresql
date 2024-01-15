package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected server error",
	}
}
func NewStatusInternalServerError(message string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func New(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
