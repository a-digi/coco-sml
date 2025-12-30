package binary

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/a-digi/coco-sml/src/db/binary/sql"
)

// Server represents the server instance and its configuration.
type Server struct {
	DataDir string // Path to the database directory (full path to the database folder)
	started bool   // Indicates if the server has been started
}

// Start initializes the server, validates the database name, and creates the folder if necessary.
// Returns the initialized Server or an error.
func Start(dataDir string) (*Server, error) {
	base := filepath.Base(dataDir)
	if !IsValidName(base) {
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

// DatabaseMeta holds metadata about the database instance.
type DatabaseMeta struct {
	Name        string    // Database name (directory)
	Version     int       // Database version
	CreatedAt   time.Time // Creation timestamp
	UpdatedAt   time.Time // Last update timestamp
	Description string    // Optional description
	Owner       string    // Optional owner/creator
	LastAccess  time.Time // Last accessed timestamp
}

// Database returns the Database instance for the given database name (directory) if the server has been started, otherwise returns an error.
// It is a method of Server.
func (s *Server) Database(dbName string) (*Database, error) {
	if !s.started {
		return nil, ErrServerNotStarted
	}

	// Validate dbName
	if !IsValidName(dbName) {
		return nil, ErrInvalidDatabaseName
	}

	// Prüfen, ob die Datenbank existiert
	_, err := s.FindDatabase(dbName)
	if err != nil {
		return nil, fmt.Errorf("database '%s' does not exist: %w", dbName, err)
	}

	fullPath := filepath.Join(s.DataDir, dbName)
	return &Database{
		Name:      fullPath,
		StartedAt: time.Now(),
	}, nil
}

// LoadTableMeta loads the metadata for all tables from tables.meta in the database directory using the shared logic from table.go.
func (db *Database) LoadTableMeta() ([]TableMeta, error) {
	return LoadTableMeta(db.Name)
}

// Select parses the SQL string query, then executes it using the shared ExecuteSelect logic from select.go.
// It returns the result rows or an error.
func (db *Database) Select(sqlStringQuery string) ([]ResultRow, error) {
	parsedQuery, err := sql.ParseSQL(sqlStringQuery)
	if err != nil {
		return nil, err
	}
	return ExecuteSelect(parsedQuery, db.Name)
}

// SaveMeta writes the database metadata to db.meta in the database directory.
func (db *Database) SaveMeta(meta DatabaseMeta) error {
	metaPath := filepath.Join(db.Name, "db.meta")
	file, err := os.Create(metaPath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	return enc.Encode(meta)
}

// LoadMeta reads the database metadata from db.meta in the database directory.
func (db *Database) LoadMeta() (DatabaseMeta, error) {
	metaPath := filepath.Join(db.Name, "db.meta")
	file, err := os.Open(metaPath)
	if err != nil {
		return DatabaseMeta{}, err
	}
	defer file.Close()
	var meta DatabaseMeta
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&meta); err != nil {
		return DatabaseMeta{}, err
	}
	return meta, nil
}

// DatabaseCreate legt eine neue Datenbank mit dem angegebenen Namen und optionaler Beschreibung an.
// Die Metadaten werden in der Binärdatei "databases" im DataDir gespeichert.
func (s *Server) DatabaseCreate(name string, description ...string) error {
	if !s.started {
		return ErrServerNotStarted
	}
	if !IsValidName(name) {
		return ErrInvalidDatabaseName
	}

	// Existenz prüfen mit FindDatabase
	_, err := s.FindDatabase(name)
	if err == nil {
		return fmt.Errorf("database '%s' already exists", name)
	}
	// Nur fortfahren, wenn der Fehler bedeutet, dass die DB nicht existiert
	if err != nil && err.Error() != fmt.Sprintf("database '%s' not found", name) {
		return err
	}

	dbMetaPath := filepath.Join(s.DataDir, "databases")
	if _, err := os.Stat(dbMetaPath); os.IsNotExist(err) {
		file, err := os.Create(dbMetaPath)
		if err != nil {
			return fmt.Errorf("could not create databases file: %w", err)
		}
		file.Close()
		fmt.Println("[DEBUG] Created new databases file:", dbMetaPath)
	}

	var dbs []DatabaseMeta
	if file, err := os.Open(dbMetaPath); err == nil {
		defer file.Close()
		dec := gob.NewDecoder(file)
		_ = dec.Decode(&dbs)
	}

	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}

	newDB := DatabaseMeta{
		Name:        name,
		Version:     1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: desc,
	}
	dbs = append(dbs, newDB)

	file, err := os.Create(dbMetaPath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	return enc.Encode(dbs)
}

// FindDatabase sucht eine Datenbank anhand des Namens in der zentralen "databases"-Datei und gibt deren Metadaten zurück.
func (s *Server) FindDatabase(name string) (DatabaseMeta, error) {
	if !s.started {
		return DatabaseMeta{}, ErrServerNotStarted
	}
	if !IsValidName(name) {
		return DatabaseMeta{}, ErrInvalidDatabaseName
	}

	dbMetaPath := filepath.Join(s.DataDir, "databases")
	file, err := os.Open(dbMetaPath)
	if err != nil {
		return DatabaseMeta{}, fmt.Errorf("could not open databases file: %w", err)
	}
	defer file.Close()

	var dbs []DatabaseMeta
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&dbs); err != nil {
		return DatabaseMeta{}, fmt.Errorf("could not decode databases file: %w", err)
	}

	for _, db := range dbs {
		if db.Name == name {
			return db, nil
		}
	}

	return DatabaseMeta{}, fmt.Errorf("database '%s' not found", name)
}

// ErrInvalidDatabaseName is returned when the database name contains invalid characters.
var ErrInvalidDatabaseName = os.ErrInvalid

// ErrServerNotStarted is returned when Database is called before Start.
var ErrServerNotStarted = os.ErrPermission
