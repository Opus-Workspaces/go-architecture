package models

import "fmt"

type ResponseError struct {
	StatusCode int         `json:"status_code"`
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type UnmarshalError struct {
	Payload    []byte `json:"payload"`
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	Err        error  `json:"error"`
	Data       string `json:"data"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("StatusCode: %d, Type: %s, Message: %s", e.StatusCode, e.Type, e.Message)
}

func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("Payload: %s, StatusCode: %d, Type: %s, Message: %s, Error: %s, Data: %s", e.Payload, e.StatusCode, e.Type, e.Message, e.Err, e.Data)
}
