package cphalo

import (
	"fmt"
	"net/http"
)

var _ CPHaloResponseError = &ResponseErrorGeneral{}
var _ CPHaloResponseError = &ResponseError400{}
var _ CPHaloResponseError = &ResponseError404{}
var _ CPHaloResponseError = &ResponseError422{}
var _ CPHaloResponseError = &ResponseError429{}
var _ CPHaloResponseError = &ResponseError500{}

type CPHaloResponseError interface {
	Error() string
	GetStatusCode() int
}

type ResponseErrorGeneral struct {
	Message    string `json:"expectedErr"`
	StatusCode int
}

func (e ResponseErrorGeneral) Error() string {
	return e.Message
}

func (e ResponseErrorGeneral) GetStatusCode() int {
	return e.StatusCode
}

type ResponseError400 struct {
	Message    string
	StatusCode int
}

func (e ResponseError400) Error() string {
	return fmt.Sprintf("request failed with %d: %s", e.StatusCode, e.Message)
}

func (e ResponseError400) GetStatusCode() int {
	return e.StatusCode
}

type ResponseError404 struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}

func (e ResponseError404) Error() string {
	return fmt.Sprintf("resource %s with %s=%s not found", e.Resource, e.Field, e.Value)
}

func (e ResponseError404) GetStatusCode() int {
	return http.StatusNotFound
}

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

func (e ResponseError422) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}

type ResponseError429 struct{}

func (e ResponseError429) Error() string {
	return http.StatusText(e.GetStatusCode())
}

func (e ResponseError429) GetStatusCode() int {
	return http.StatusTooManyRequests
}

type ResponseError500 struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (e ResponseError500) Error() string {
	return fmt.Sprintf("server failed with code %d and expectedErr: %s", e.StatusCode, e.Message)
}

func (e ResponseError500) GetStatusCode() int {
	return e.StatusCode
}
