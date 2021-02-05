package helper

import "strings"

// Response untuk object response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// EmptyObj jika nilai tidak boleh null/nil
type EmptyObj struct {
}

// BuildResponse untuk response sukses
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// BuildErrorResponse untuk response gagal
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status: false,
		Message: message,
		Error: splittedError,
		Data: data,
	}
	return res
}
