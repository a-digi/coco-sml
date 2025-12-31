package di

import (
	"github.com/a-digi/coco-sml/src/api/auth/encrypt"
)

type ServiceBag struct {
	BcryptManager   *encrypt.BcryptManager
	RegistryManager any // late binding to avoid import cycle
}

// NewServiceBag initializes and returns a ServiceBag with all dependencies wired up.
func NewServiceBag() *ServiceBag {
	return &ServiceBag{
		BcryptManager:   encrypt.CreateBcryptManager(),
		RegistryManager: nil, // set after construction
	}
}

func (sb *ServiceBag) SetRegistryManager(rm any) {
	sb.RegistryManager = rm
}
