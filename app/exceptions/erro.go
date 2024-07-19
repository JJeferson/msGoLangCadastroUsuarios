package exceptions

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func NewAppError(message string, err error) *AppError {
	return &AppError{Message: message, Err: err}
}

func HandleError(w http.ResponseWriter, message string, err error) {
	appErr := NewAppError(message, err)
	http.Error(w, appErr.Error(), http.StatusInternalServerError)
}
