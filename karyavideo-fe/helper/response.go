package helper

import "strings"

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Errors interface {} `json:"errors"`
	Data interface {} `json:"data"`
}

type EmptyObj struct {}

func BuildResponse (status bool, message string, data interface{}) Response{
	res:= Response {
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}
	return res
}

func BuildErrorResponse (status bool, message string, data interface{}) Response{
	SplitErrors := strings.Split(message,"")
	res:= Response {
		Status:  status,
		Message: message,
		Errors:  SplitErrors,
		Data:    data,
	}
	return res
}
