package response

import (
	"encoding/json"
)

// StatusType is the status of a response
// success = success, failed = error

type StatusType string

const (
	StatusSuccess StatusType = "success"
	StatusFailed  StatusType = "failed"
)

// Response is the base struct for all API responses
// It is serialized as JSON

type Response struct {
	Status  StatusType   `json:"status"`
	Message interface{}  `json:"message"`
}

// Success represents a successful response

type Success struct {
	Response
}

// Error represents an error response

type Error struct {
	Response
}

// SuccessResponse creates a Success response and returns it as JSON bytes
func SuccessResponse(message interface{}) []byte {
	resp := Success{
		Response: Response{
			Status:  StatusSuccess,
			Message: message,
		},
	}
	jsonBytes, _ := json.Marshal(resp)
	return jsonBytes
}

// ErrorResponse creates an Error response and returns it as JSON bytes
func ErrorResponse(message interface{}) []byte {
	resp := Error{
		Response: Response{
			Status:  StatusFailed,
			Message: message,
		},
	}
	jsonBytes, _ := json.Marshal(resp)
	return jsonBytes
}
