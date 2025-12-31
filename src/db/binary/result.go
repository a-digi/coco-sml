package binary

import (
	"encoding/json"
)

// StatusType is an enum for Result status.
type StatusType string

const (
	StatusSuccess StatusType = "success"
	StatusError   StatusType = "error"
)

// Result represents the outcome of a database action, including the result data, the action performed, the status, and the time it took in milliseconds.
type Result struct {
	Results []interface{} `json:"results"` // The result data (array of results)
	Action  string        `json:"action"`  // The action performed (e.g., "select", "insert")
	Status  StatusType    `json:"status"`  // The status of the action (success or error)
	TimeMs  float64       `json:"timeMs"`  // The time it took in milliseconds
}

// CreateResult is a factory function to create a Result instance.
func CreateResult(results []interface{}, action string, status StatusType, timeMs float64) Result {
	return Result{
		Results: results,
		Action:  action,
		Status:  status,
		TimeMs:  timeMs,
	}
}

// ResultSuccess creates a Result with status set to success.
func ResultSuccess(results []interface{}, action string, timeMs float64) Result {
	return Result{
		Results: results,
		Action:  action,
		Status:  StatusSuccess,
		TimeMs:  timeMs,
	}
}

// ResultError creates a Result with status set to error and an empty array as results.
func ResultError(action string, timeMs float64) Result {
	return Result{
		Results: []interface{}{},
		Action:  action,
		Status:  StatusError,
		TimeMs:  timeMs,
	}
}

// ConvertJSON returns the JSON encoding of the Result or an error if encoding fails.
func (r Result) ConvertJSON() ([]byte, error) {
	return json.Marshal(r)
}