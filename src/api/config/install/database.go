package install

import (
	"log"
	"github.com/a-digi/coco-sml/src/db/binary"
)

// SetupInternalDatabase checks if the 'internal' database exists and creates it if not.
func SetupInternalDatabase(server *binary.Server) {
	_, err := server.FindDatabase("internal")
	if err != nil {
		errCreate := server.DatabaseCreate("internal", "Internal system database")
		if errCreate != nil {
			log.Printf("[SetupApi] Failed to create database 'internal': %v\n", errCreate)
		}
	}
}

