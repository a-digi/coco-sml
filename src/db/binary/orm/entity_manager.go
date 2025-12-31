package orm

import (
	"github.com/a-digi/coco-sml/src/db/binary"
)

// EntityManager is responsible for managing entities within a specific database instance.
type EntityManager struct {
	Database *binary.Database
}

// NewEntityManager creates and returns a new EntityManager for the given database.
func NewEntityManager(db *binary.Database) *EntityManager {
	return &EntityManager{
		Database: db,
	}
}
