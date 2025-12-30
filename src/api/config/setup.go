package config

import (
	"github.com/a-digi/coco-sml/src/db/binary"
	"log"
)

// SetupApi initializes the API setup with a given binary.Server
func SetupApi(server *binary.Server) {
	log.Println("[SetupApi] Checking for database 'internal' ...")
	_, err := server.FindDatabase("internal")
	if err != nil {
		log.Println("[SetupApi] Database 'internal' does not exist. Creating...")
		errCreate := server.DatabaseCreate("internal", "Internal system database")
		if errCreate != nil {
			log.Printf("[SetupApi] Failed to create database 'internal': %v\n", errCreate)
		} else {
			log.Println("[SetupApi] Database 'internal' created successfully.")
		}
	} else {
		log.Println("[SetupApi] Database 'internal' already exists.")
	}
}
