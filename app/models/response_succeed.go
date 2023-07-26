package models

import (
	"fmt"
)

type ResponseSucceed struct {
	StatusCode int         `json:"status_code"`
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type UnmarshalSucceed struct {
	Payload    []byte      `json:"payload"`
	StatusCode int         `json:"status_code"`
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (s *ResponseSucceed) Error() string {
	return fmt.Sprintf("StatusCode: %d, Type: %s, Message: %s", s.StatusCode, s.Type, s.Message)
}

func (s *UnmarshalSucceed) Error() string {
	return fmt.Sprintf("Payload: %s, StatusCode: %d, Type: %s, Message: %s, Data: %s", s.Payload, s.StatusCode, s.Type, s.Message, s.Data)
}
