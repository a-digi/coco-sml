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

	// Define the user table schema
	table := binary.Table{
		Name:        "user",
		Description: "User table for authentication",
		Fields: []binary.Field{
			{Name: "id", DataType: binary.UUIDType},
			{Name: "username", DataType: binary.StringType, MinLength: 1, MaxLength: 255, Required: true},
			{Name: "password", DataType: binary.StringType, MinLength: 1, MaxLength: 255, Required: true},
			{Name: "createdAt", DataType: binary.DateTimeType, Required: true, DefaultValue: "now"},
			{Name: "isActive", DataType: binary.BoolType, DefaultValue: true},
		},
	}

	// Try to create the table (ignore error if it already exists)
	err = binary.CreateTable(db.Name, table)
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

	// Define the admin_user table schema
	table := binary.Table{
		Name:        "root_user",
		Description: "Root user table for system administrators",
		Fields: []binary.Field{
			{Name: "id", DataType: binary.UUIDType},
			{Name: "username", DataType: binary.StringType, MinLength: 1, MaxLength: 255, Required: true},
			{Name: "password", DataType: binary.StringType, MinLength: 1, MaxLength: 255, Required: true},
			{Name: "email", DataType: binary.StringType, MinLength: 1, MaxLength: 255, Required: true},
			{Name: "createdAt", DataType: binary.DateTimeType, DefaultValue: "now"},
			{Name: "isActive", DataType: binary.BoolType, DefaultValue: true},
		},
	}

	// Try to create the table (ignore error if it already exists)
	err = binary.CreateTable(db.Name, table)
	if err != nil {
		if err.Error() == "table already exists" {
			return
		}
		log.Printf("[SetupAdminUserTable] Failed to create 'admin_user' table: %v\n", err)
	}
}
