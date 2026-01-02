package event

import (
    "encoding/binary"
    "fmt"
    "io"
    "os"
    "sync"
    "time"
)

// InsertEvent represents a single insert event with metadata
// Payload must be a raw SQL string (e.g., an INSERT statement or SQL row data)
type InsertEvent struct {
    Table     string    // Table name
    Payload   string    // Raw SQL string for the insert operation
    Timestamp time.Time // Event timestamp
}

// Internal event queue for insert events
var (
    eventQueue chan InsertEvent
    once       sync.Once
    fileWriteMutex sync.Mutex // Ensures race condition safety for file writes
)

// InitializeQueue sets up the internal event queue
func InitializeQueue(bufferSize int) {
    once.Do(func() {
        eventQueue = make(chan InsertEvent, bufferSize)
    })
}

// EnqueueInsert validates and adds an event to the queue
// The payload must be a valid SQL string
func EnqueueInsert(table string, payload string) error {
    // TODO: Add SQL string validation if needed
    event := InsertEvent{
        Table:     table,
        Payload:   payload, // Must be a SQL string
        Timestamp: time.Now(),
    }
    eventQueue <- event
    return nil
}

// Config holds configuration for event processing
// Add more fields as needed for your environment
// EventFolderPath is used instead of EventFilePath
type Config struct {
    EventFolderPath string // Path to the folder where event files are stored
}

// WriteEventToFile appends an event to the table-specific event file in a length-prefixed binary format
// Format: [tableLen][table][payloadLen][payload][timestamp]
// Each table has its own event file: <table>_events
func WriteEventToFile(event InsertEvent, cfg Config) error {
    eventFilePath := fmt.Sprintf("%s/%s_events", cfg.EventFolderPath, event.Table)
    fileWriteMutex.Lock()
    defer fileWriteMutex.Unlock()
    f, err := os.OpenFile(eventFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    tableBytes := []byte(event.Table)
    payloadBytes := []byte(event.Payload)
    tableLen := uint32(len(tableBytes))
    payloadLen := uint32(len(payloadBytes))
    timestamp := event.Timestamp.UnixNano()

    // Write table length
    if err := binary.Write(f, binary.LittleEndian, tableLen); err != nil {
        return err
    }
    // Write table bytes
    if _, err := f.Write(tableBytes); err != nil {
        return err
    }
    // Write payload length
    if err := binary.Write(f, binary.LittleEndian, payloadLen); err != nil {
        return err
    }
    // Write payload bytes
    if _, err := f.Write(payloadBytes); err != nil {
        return err
    }
    // Write timestamp
    if err := binary.Write(f, binary.LittleEndian, timestamp); err != nil {
        return err
    }
    return f.Sync()
}

// ReadNextEventFromFile reads the next InsertEvent from the event file (for demonstration/testing)
func ReadNextEventFromFile(r io.Reader) (*InsertEvent, error) {

    var tableLen uint32
    if err := binary.Read(r, binary.LittleEndian, &tableLen); err != nil {
        return nil, err
    }

    tableBytes := make([]byte, tableLen)
    if _, err := io.ReadFull(r, tableBytes); err != nil {
        return nil, err
    }

    var payloadLen uint32
    if err := binary.Read(r, binary.LittleEndian, &payloadLen); err != nil {
        return nil, err
    }

    payloadBytes := make([]byte, payloadLen)
    if _, err := io.ReadFull(r, payloadBytes); err != nil {
        return nil, err
    }
    var timestamp int64
    if err := binary.Read(r, binary.LittleEndian, &timestamp); err != nil {
        return nil, err
    }
    return &InsertEvent{
        Table:     string(tableBytes),
        Payload:   string(payloadBytes),
        Timestamp: time.Unix(0, timestamp),
    }, nil
}

// StartEventConsumer launches a goroutine to process events from the queue
// Receives configuration struct for dynamic event file path and other options
func StartEventConsumer(cfg Config) {
    go func() {
        for event := range eventQueue {
            // Write event to file using dynamic path from config
            err := WriteEventToFile(event, cfg)
            if err != nil {
                // Log error, optionally retry or handle
                continue
            }
            // TODO: Add logic to process event and insert into DB
            // e.g., call InsertToTable(event)
        }
    }()
}

// Offset file helpers for tracking processed events
// Each table has its own offset file: <table>_events.offset

// GetOffsetFilePath returns the offset file path for a table
func GetOffsetFilePath(cfg Config, table string) string {
    return fmt.Sprintf("%s/%s_events.offset", cfg.EventFolderPath, table)
}

// ReadLastProcessedOffset reads the last processed byte offset from the offset file
func ReadLastProcessedOffset(cfg Config, table string) (int64, error) {
    offsetFile := GetOffsetFilePath(cfg, table)
    f, err := os.Open(offsetFile)
    if err != nil {
        if os.IsNotExist(err) {
            return 0, nil // If file doesn't exist, start from 0
        }
        return 0, err
    }
    defer f.Close()
    var offset int64
    err = binary.Read(f, binary.LittleEndian, &offset)
    if err != nil && err != io.EOF {
        return 0, err
    }

    return offset, nil
}

// WriteLastProcessedOffset writes the last processed byte offset to the offset file
func WriteLastProcessedOffset(cfg Config, table string, offset int64) error {
    offsetFile := GetOffsetFilePath(cfg, table)
    f, err := os.OpenFile(offsetFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    return binary.Write(f, binary.LittleEndian, offset)
}

// ProcessEventsFromFile processes unprocessed events from the event file using offset tracking
func ProcessEventsFromFile(cfg Config, table string, processFunc func(event *InsertEvent) error) error {
    eventFilePath := fmt.Sprintf("%s/%s_events", cfg.EventFolderPath, table)
    offset, err := ReadLastProcessedOffset(cfg, table)
    if err != nil {
        return err
    }
    f, err := os.Open(eventFilePath)
    if err != nil {
        return err
    }
    defer f.Close()
    // Seek to last processed offset
    _, err = f.Seek(offset, io.SeekStart)

    if err != nil {
        return err
    }

    for {
        currOffset, _ := f.Seek(0, io.SeekCurrent)
        event, err := ReadNextEventFromFile(f)
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        // Process the event
        if err := processFunc(event); err != nil {
            return err
        }
        // Update offset after successful processing
        nextOffset, _ := f.Seek(0, io.SeekCurrent)
        if err := WriteLastProcessedOffset(cfg, table, nextOffset); err != nil {
            return err
        }
    }

    return nil
}
