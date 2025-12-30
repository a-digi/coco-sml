package internaldb

import (
	"encoding/binary"
	"errors"
	"os"
)

const (
	// Adjust these sizes as needed for your actual data
	MaxConditions = 8
	MaxFieldLen   = 32
	MaxValueLen   = 256
	EntrySize     = MaxConditions*(MaxFieldLen+2+MaxValueLen) + 1 + MaxValueLen // rough estimate
)

// Condition represents a single search/filter condition for an entry.
// Field: the name of the field to compare
// Operator: the comparison operator (e.g., "=", "!=", ">", "<")
// Value: the value to compare against
// For extensibility, you can add more fields as needed.
type Condition struct {
	Field    string
	Operator string
	Value    string
}

// Entry represents a record in the binary search engine.
// Conditions: array of conditions that uniquely identify the entry
// Value: associated value for the entry
// Deleted: flag for soft deletion (true if entry is logically deleted)
type Entry struct {
	Conditions []Condition
	Value      string
	Deleted    bool
}

// ErrEntryNotFound is returned when no matching entry is found in the file.
var ErrEntryNotFound = errors.New("entry not found")

// compareConditions checks if two arrays of conditions are equal (for binary search).
// Only supports equality comparison for now.
func compareConditions(a, b []Condition) bool {

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// BinarySearch performs a binary search for the given conditions in a sorted slice of entries.
// Returns the index and a pointer to the Entry, or nil if not found or deleted.
// The entries slice must be sorted by Conditions in ascending order (lexicographically).
// Returns -1 and nil if the conditions are not found or the entry is deleted.
func BinarySearch(entries []Entry, conditions []Condition) (int, *Entry) {

	if len(entries) == 0 {
		return -1, nil
	}

	low, high := 0, len(entries)-1

	for low <= high {
		mid := (low + high) / 2
		if compareConditions(entries[mid].Conditions, conditions) {
			if entries[mid].Deleted {
				return mid, nil
			}
			return mid, &entries[mid]
		} else if lessConditions(entries[mid].Conditions, conditions) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, nil
}

// lessConditions compares two arrays of conditions lexicographically for sorting.
// Returns true if a < b.
func lessConditions(a, b []Condition) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i].Field < b[i].Field {
			return true
		} else if a[i].Field > b[i].Field {
			return false
		}
		if a[i].Operator < b[i].Operator {
			return true
		} else if a[i].Operator > b[i].Operator {
			return false
		}
		if a[i].Value < b[i].Value {
			return true
		} else if a[i].Value > b[i].Value {
			return false
		}
	}
	return len(a) < len(b)
}

// BinarySearchFile performs a binary search for the given conditions in a binary file of entries.
// Assumes fixed-size records for demonstration. Returns the entry and its index, or error.
func BinarySearchFile(filePath string, conditions []Condition) (int64, *Entry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return -1, nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return -1, nil, err
	}
	totalEntries := fileInfo.Size() / int64(EntrySize)
	low, high := int64(0), totalEntries-1

	for low <= high {
		mid := (low + high) / 2
		pos := mid * int64(EntrySize)
		entry, err := readEntryAt(file, pos)
		if err != nil {
			return -1, nil, err
		}
		if compareConditions(entry.Conditions, conditions) {
			if entry.Deleted {
				return mid, nil, ErrEntryNotFound
			}
			return mid, entry, nil
		} else if lessConditions(entry.Conditions, conditions) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, nil, ErrEntryNotFound
}

// readEntryAt reads a single entry from the file at the given offset.
// Assumes fixed-size records for demonstration.
func readEntryAt(file *os.File, offset int64) (*Entry, error) {
	buf := make([]byte, EntrySize)
	_, err := file.ReadAt(buf, offset)
	if err != nil {
		return nil, err
	}
	// TODO: Implement proper decoding from buf to Entry
	// For demonstration, return a zero Entry
	return &Entry{}, nil
}
