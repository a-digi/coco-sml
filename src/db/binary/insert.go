package binary

import (
	"github.com/a-digi/coco-sml/src/db/binary/sql"
	"github.com/a-digi/coco-sml/src/db/binary/insert"
)

/*
Event Triggering and Processing Overview

Event-sourced data insertion and processing will be handled internally using Go channels and goroutines. No REST endpoints will be exposed for event listing, consumption, or status reporting. All operations will be triggered and managed via Go's concurrency primitives for optimal performance and reliability.

- Event Creation (Internal Queue):
    - Data to be inserted is validated and placed into an internal queue (e.g., Go channel).
    - The queue acts as the buffer for incoming insert requests before event sourcing.
    - Each item in the queue is written as an event to an append-only file with a timestamp.
    - This step is triggered internally (e.g., via CLI, scheduled job, or direct function call).

- Event Consumption (Backend Consumers):
    - Dedicated backend consumers (goroutines) monitor the event file for new events.
    - Consumers read, validate, and insert events into the database tables.
    - Communication and coordination are managed via Go channels and internal queue mechanisms.

- Status and Monitoring:
    - Status of event processing and table insertion is tracked internally.
    - Monitoring and auditing can be implemented using logs, metrics, or internal status objects, not via HTTP endpoints.

Note: All event flow, consumption, and status tracking are handled within the Go application using internal queues (channels) and backend consumers (goroutines). No REST API endpoints are provided for these operations.
*/

// Insert parses an SQL string, creates an event, and returns a result.Result (success or error).
func Insert(sqlString string) Result {
	parsed, err := sql.ParseInsertSQL(sqlString)
	if err != nil {
		return ResultError("insert", nil)
	}

	err = insert.InsertToTable(parsed)
	if err != nil {
		return ResultError("insert", nil)
	}

	return ResultSuccess([]interface{}{}, "insert", nil)
}
