package orm

import (
	"errors"
	"github.com/a-digi/coco-sml/src/db/binary"
	"log"
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

// ListTables returns all tables (with full metadata) in the database managed by this EntityManager.
func (em *EntityManager) ListTables() ([]binary.Table, error) {
	log.Println("[EntityManager] ListTables called")
	if em.Database == nil {
		log.Println("[EntityManager] Database is not initialized")
		return nil, errors.New("database is not initialized")
	}
	log.Printf("[EntityManager] Calling ListTables on Database: %s\n", em.Database.Name)
	tables, err := em.Database.ListTables()
	if err != nil {
		log.Printf("[EntityManager] Database.ListTables error: %v\n", err)
	}
	return tables, err
}
