package cphalo

import (
	"fmt"
	"net/http"
)

var _ ResponseError = &ResponseErrorGeneral{}
var _ ResponseError = &ResponseError400{}
var _ ResponseError = &ResponseError404{}
var _ ResponseError = &ResponseError422{}
var _ ResponseError = &ResponseError429{}
var _ ResponseError = &ResponseError500{}

// ResponseError is a interface for all response errors.
type ResponseError interface {
	Error() string
	GetStatusCode() int
}

// ResponseErrorGeneral is a representation of general error.
type ResponseErrorGeneral struct {
	Message    string `json:"expectedErr"`
	StatusCode int
}

func (e ResponseErrorGeneral) Error() string {
	return e.Message
}

// GetStatusCode returns response status code.
func (e ResponseErrorGeneral) GetStatusCode() int {
	return e.StatusCode
}

// ResponseError400 is a representation of 400 error.
type ResponseError400 struct {
	Message    string
	StatusCode int
}

func (e ResponseError400) Error() string {
	return fmt.Sprintf("request failed with %d: %s", e.StatusCode, e.Message)
}

// GetStatusCode returns response status code.
func (e ResponseError400) GetStatusCode() int {
	return e.StatusCode
}

// ResponseError404 is a representation of 404 error.
type ResponseError404 struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}

func (e ResponseError404) Error() string {
	return fmt.Sprintf("resource %s with %s=%s not found", e.Resource, e.Field, e.Value)
}

// GetStatusCode returns response status code.
func (e ResponseError404) GetStatusCode() int {
	return http.StatusNotFound
}

// ResponseError422 is a representation of 422 error.
type ResponseError422 struct {
	Message string `json:"message"`
	Errors  []struct {
		Field   string `json:"field"`
		Value   string `json:"value"`
		Code    string `json:"code"`
		Details string `json:"details"`
	} `json:"errors"`
}

func (e ResponseError422) Error() string {
	msg := fmt.Sprintf("validation failed for %d fields with msg: %s", len(e.Errors), e.Message)

	for _, e := range e.Errors {
		msg = msg + fmt.Sprintf("\n- %s %s", e.Field, e.Code)
	}

	return msg
}

// GetStatusCode returns response status code.
func (e ResponseError422) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}

// ResponseError429 is a representation of 429 error.
type ResponseError429 struct{}

func (e ResponseError429) Error() string {
	return http.StatusText(e.GetStatusCode())
}

// GetStatusCode returns response status code.
func (e ResponseError429) GetStatusCode() int {
	return http.StatusTooManyRequests
}

// ResponseError500 is a representation of 500 error.
type ResponseError500 struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}

func (e ResponseError500) Error() string {
	return fmt.Sprintf("server failed with code %d and expectedErr: %s", e.StatusCode, e.Message)
}

// GetStatusCode returns response status code.
func (e ResponseError500) GetStatusCode() int {
	return e.StatusCode
}
