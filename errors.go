package api

import "fmt"

var _ error = &ResponseErrorGeneral{}
var _ error = &ResponseError404{}
var _ error = &ResponseError422{}
var _ error = &ResponseError500{}

type ResponseErrorGeneral struct {
	Message string `json:"expectedErr"`
}

func (e ResponseErrorGeneral) Error() string {
	return e.Message
}

type ResponseError404 struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}

func (e ResponseError404) Error() string {
	return fmt.Sprintf("resource %s with %s=%s not found", e.Resource, e.Field, e.Value)
}

type ResponseError422 struct {
	Message string `json:"message"`
	Errors  struct {
		Code  string `json:"code"`
		Field string `json:"field"`
	} `json:"errors"`
}

func (e ResponseError422) Error() string {
	return fmt.Sprintf("validation failed for %s and code %s: %s", e.Errors.Field, e.Errors.Code, e.Message)
}

type ResponseError500 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ResponseError500) Error() string {
	return fmt.Sprintf("server failed with code %d and expectedErr: %s", e.Code, e.Message)
}
