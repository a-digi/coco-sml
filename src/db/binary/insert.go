package binary

// InsertEntry inserts or updates an entry in a sorted slice of entries.
// If the key exists and is deleted, it will be restored and updated.
// Returns the updated slice.
func InsertEntry(entries []Entry, key, value string) []Entry {
	idx, entry := BinarySearch(entries, key)
	if entry != nil {
		entries[idx].Value = value
		entries[idx].Deleted = false
		return entries
	}
	entries = append(entries, Entry{})
	i := len(entries) - 2
	for i >= 0 && entries[i].Key > key {
		entries[i+1] = entries[i]
		i--
	}
	entries[i+1] = Entry{Key: key, Value: value, Deleted: false}
	return entries
}

// DeleteEntry marks an entry as deleted by key in a sorted slice of entries.
// Returns the updated slice. If the key does not exist, the slice is unchanged.
func DeleteEntry(entries []Entry, key string) []Entry {
	idx, entry := BinarySearch(entries, key)
	if entry != nil {
		entries[idx].Deleted = true
	}
	return entries
}

