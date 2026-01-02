package install

import (
	"log"
	"github.com/a-digi/coco-sml/src/db/binary"
)

// SetupUserTable creates a 'user' table with username and password fields if it does not exist in the 'internal' database.
func SetupUserTable(server *binary.Server) {
	// Get the internal database
	db, err := server.Database("internal")
	if err != nil {
		log.Printf("[SetupUserTable] Could not access 'internal' database: %v\n", err)
		return
	}

	// Define the user table schema as a SQL string
	sql := `CREATE TABLE user (
		id UUID,
		username VARCHAR,
		password VARCHAR,
		createdAt DATETIME,
		isActive BOOL
	);`

	// Try to create the table using db.CreateTable (ignore error if it already exists)
	err = db.CreateTable(sql)
	if err != nil {
		if err.Error() == "table already exists" {
			return
		}
		log.Printf("[SetupUserTable] Failed to create 'user' table: %v\n", err)
	}
}

// SetupRootUserTable creates an 'root_user' table with admin-specific fields if it does not exist in the 'internal' database.
func SetupRootUserTable(server *binary.Server) {
	// Get the internal database
	db, err := server.Database("internal")
	if err != nil {
		log.Printf("[SetupRootUserTable] Could not access 'internal' database: %v\n", err)
		return
	}

	// Define the admin_user table schema as a SQL string
	sql := `CREATE TABLE root_user (
		id UUID,
		username VARCHAR,
		password VARCHAR,
		email VARCHAR,
		createdAt DATETIME,
		isActive BOOL
	);`

	// Try to create the table using db.CreateTable (ignore error if it already exists)
	err = db.CreateTable(sql)
	if err != nil {
		if err.Error() == "table already exists" {
			return
		}
		log.Printf("[SetupAdminUserTable] Failed to create 'root_user' table: %v\n", err)
	}
}
