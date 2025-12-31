package orm

import (
	"errors"
	"sync"
	"github.com/a-digi/coco-sml/src/db/binary"
	"github.com/a-digi/coco-sml/src/db/binary/orm"
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
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.Server == nil {
		return nil, errors.New("server is not initialized")
	}

	if em, ok := rm.cache[databaseName]; ok {
		return em, nil
	}

	db, err := rm.Server.Database(databaseName)
	if err != nil {
		return nil, err
	}

	em := NewEntityManager(db)
	rm.cache[databaseName] = em

	return em, nil
}

func (rm *RegistryManager) GetManager(databaseName string) (*EntityManager, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if em, ok := rm.cache[databaseName]; ok {
		return em, nil
	}

	em, err := rm.Create(databaseName)

	if err != nil {
		return nil, err
	}

	return em, nil
}
