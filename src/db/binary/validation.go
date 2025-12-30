package binary

import (
	"regexp"
)

// IsValidDatabaseName prüft, ob der Datenbankname gültig ist (nur Buchstaben und Unterstriche).
func IsValidName(name string) bool {
	validName := regexp.MustCompile(`^[A-Za-z_]+$`)
	return validName.MatchString(name)
}

