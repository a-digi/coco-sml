package internaldb

// DeleteEntry marks an entry as deleted by key in a sorted slice of entries.
// Returns the updated slice. If the key does not exist, the slice is unchanged.
func DeleteEntry(entries []Entry, key string) []Entry {
	idx, entry := BinarySearch(entries, key)

	if entry != nil {
		entries[idx].Deleted = true
	}

	return entries
}

