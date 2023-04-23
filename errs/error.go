package errs

import "net/http"

type AppError struct {
	Code int `json:",omitempty"` // 在发送response的时候忽略
	Message string `json:"message"`
}

func NewInterError(msg string) *AppError {
	return &AppError{
		Code: http.StatusInternalServerError,
		Message: msg,
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Code: http.StatusNotFound,
		Message: msg,
	}
}

func NewValidateError(msg string) *AppError {
	return &AppError{
		Code: http.StatusUnprocessableEntity,
		Message: msg,
	}
}

func (e *AppError) AsMessage() *AppError  {
	return &AppError{
		Message: e.Message,
	}
}