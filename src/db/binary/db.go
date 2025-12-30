package binary

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	table "github.com/a-digi/coco-sml/src/db/internal/table"
)

// Server represents the server instance and its configuration.
type Server struct {
	DataDir string // Path to the database directory (full path to the database folder)
	started bool   // Indicates if the server has been started
}

// Start initializes the server, validates the database name, and creates the folder if necessary.
// Returns the initialized Server or an error.
func Start(dataDir string) (*Server, error) {
	validName := regexp.MustCompile(`^[A-Za-z_]+$`)

	if !validName.MatchString(dataDir) {
		return nil, ErrInvalidDatabaseName
	}

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	return &Server{
		DataDir: dataDir,
		started: true,
	}, nil
}

// Database represents the database instance and its configuration.
type Database struct {
	Name      string    // Name of the database (full path to the database folder)
	StartedAt time.Time // Timestamp when the database was started
}

// Database returns the Database instance if the server has been started, otherwise returns an error.
// It is a method of Server.
func (s *Server) Database() (*Database, error) {

	if !s.started {
		return nil, ErrServerNotStarted
	}

	return &Database{
		Name:      s.DataDir,
		StartedAt: time.Now(),
	}, nil
}

// LoadTableMeta loads the metadata for all tables from tables.meta in the database directory using the shared logic from table.go.
func (db *Database) LoadTableMeta() ([]table.TableMeta, error) {
	return table.LoadTableMeta(db.Name)
}

// ErrInvalidDatabaseName is returned when the database name contains invalid characters.
var ErrInvalidDatabaseName = os.ErrInvalid

// ErrServerNotStarted is returned when Database is called before Start.
var ErrServerNotStarted = os.ErrPermission
