package error

import (
	"errors"
	"net/http"
)

var (
	RecordNotFound = errors.New("record not found")
)

type TemplateError struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	Err     error  `json:"error"`
}

type errorResponse struct {
	Msg string `json:"msg"`
}

func (e TemplateError) Error() string {
	return e.Err.Error()
}

func NewRecordNotFoundError(msg string, err error) TemplateError {
	return TemplateError{
		Status:  http.StatusNotFound,
		Message: msg,
		Err:     err,
	}
}

func NewUnespectedError(msg string, err error) TemplateError {
	return TemplateError{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Err:     err,
	}
}

func NewBadRequestError(msg string, err error) TemplateError {
	return TemplateError{
		Status:  http.StatusBadRequest,
		Message: msg,
		Err:     err,
	}
}

func NewValidationError(msg string, err error) TemplateError {
	return TemplateError{
		Status:  http.StatusBadRequest,
		Message: msg,
		Err:     err,
	}
}

func NewErrorResponse(message string) errorResponse {
	return errorResponse{
		Msg: message,
	}
}
