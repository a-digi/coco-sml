package orm

import (
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

// ListTables returns all tables (with full metadata) in the database managed by this EntityManager as a Result.
func (em *EntityManager) ListTablesResult() binary.Result {
	log.Println("[EntityManager] ListTablesResult called")
	if em.Database == nil {
		log.Println("[EntityManager] Database is not initialized")
		return binary.ResultError("list_tables", 0)
	}

	log.Printf("[EntityManager] Calling ListTablesResult on Database: %s\n", em.Database.Name)
	return em.Database.ListTablesResult()
}
