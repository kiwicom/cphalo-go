package api

import "fmt"

type CpHaloError struct {
	StatusCode int    `json:"statuscode"`
	StatusDesc string `json:"statusdesc"`
	Message    string `json:"errormessage"`
}

func (r *CpHaloError) Error() string {
	return fmt.Sprintf("%d %v: %v", r.StatusCode, r.StatusDesc, r.Message)
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseError404 struct {
	
}
