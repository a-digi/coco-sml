package binary

import (
	"encoding/json"
	"github.com/a-digi/coco-sml/src/db/binary/performance"
)

// StatusType is an enum for Result status.
type StatusType string

const (
	StatusSuccess StatusType = "success"
	StatusError   StatusType = "error"
)

// Result represents the outcome of a database action, including the result data, the action performed, the status, and the performance tracking information.
type Result struct {
	Results    []interface{}                `json:"results"`    // The result data (array of results)
	Action     string                       `json:"action"`     // The action performed (e.g., "select", "insert")
	Status     StatusType                   `json:"status"`     // The status of the action (success or error)
	Performance *performance.PerformanceTracker `json:"performance"` // Performance tracking information
}

// CreateResult is a factory function to create a Result instance.
func CreateResult(results []interface{}, action string, status StatusType, performance *performance.PerformanceTracker) Result {
	return Result{
		Results:    results,
		Action:     action,
		Status:     status,
		Performance: performance,
	}
}

// ResultSuccess creates a Result with status set to success.
func ResultSuccess(results []interface{}, action string, performance *performance.PerformanceTracker) Result {
	return Result{
		Results:    results,
		Action:     action,
		Status:     StatusSuccess,
		Performance: performance,
	}
}

// ResultError creates a Result with status set to error and an empty array as results.
// Accepts nil, 0, or a valid *performance.PerformanceTracker as the second argument.
func ResultError(action string, perfAny interface{}) Result {
	var perf *performance.PerformanceTracker
	switch v := perfAny.(type) {
	case nil:
		perf = performance.NewPerformanceTracker()
	case int:
		if v == 0 {
			perf = performance.NewPerformanceTracker()
		}
	case *performance.PerformanceTracker:
		perf = v
	default:
		perf = performance.NewPerformanceTracker()
	}
	return Result{
		Results:    []interface{}{},
		Action:     action,
		Status:     StatusError,
		Performance: perf,
	}
}

// ConvertJSON returns the JSON encoding of the Result or an error if encoding fails.
func (r Result) ConvertJSON() ([]byte, error) {
	return json.Marshal(r)
}
