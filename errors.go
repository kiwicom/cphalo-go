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

type ResponseError500 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ResponseError500) Error() string {
	return fmt.Sprintf("server failed with code %d and expectedErr: %s", e.Code, e.Message)
}
