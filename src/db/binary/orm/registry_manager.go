package orm

import (
	"errors"
	"log"
	"sync"

	"github.com/a-digi/coco-sml/src/db/binary"
)

type RegistryManager struct {
	Server *binary.Server
	cache  map[string]*EntityManager
	mu     sync.Mutex
}

func NewRegistryManager(server *binary.Server) *RegistryManager {
	return &RegistryManager{
		Server: server,
		cache:  make(map[string]*EntityManager),
	}
}

func (rm *RegistryManager) Create(databaseName string) (*EntityManager, error) {
	log.Printf("[RegistryManager] Create called for db: %s", databaseName)
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.Server == nil {
		log.Println("[RegistryManager] Server is not initialized")
		return nil, errors.New("server is not initialized")
	}

	if em, ok := rm.cache[databaseName]; ok {
		log.Printf("[RegistryManager] Returning cached EntityManager for db: %s", databaseName)
		return em, nil
	}

	db, err := rm.Server.Database(databaseName)
	if err != nil {
		log.Printf("[RegistryManager] Server.Database error: %v", err)
		return nil, err
	}

	em := NewEntityManager(db)
	rm.cache[databaseName] = em
	log.Printf("[RegistryManager] Created and cached new EntityManager for db: %s", databaseName)
	return em, nil
}

func (rm *RegistryManager) GetManager(databaseName string) (*EntityManager, error) {
	log.Printf("[RegistryManager] GetManager called for db: %s", databaseName)
	// Do NOT lock here, Create handles locking
	return rm.Create(databaseName)
}
