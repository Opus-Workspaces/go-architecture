package models

type ValidateType = map[string]ValidateObject

type ValidateObject struct {
	Min      uint16
	Max      uint32
	OrderBy  string
	Required bool
	Message  string
	Children map[string]ValidateObject
}
