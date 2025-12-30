package config

import (
	"github.com/a-digi/coco-sml/src/db/binary"
	"log"
)

// SetupApi initializes the API setup with a given binary.Server
func SetupApi(server *binary.Server) {
	_, err := server.FindDatabase("internal")
	if err != nil {
		errCreate := server.DatabaseCreate("internal", "Internal system database")
		if errCreate != nil {
			log.Printf("[SetupApi] Failed to create database 'internal': %v\n", errCreate)
		}
	}
}
