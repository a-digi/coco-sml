package di

import (
	"github.com/a-digi/coco-sml/src/api/auth/encrypt"
)

type ServiceBag struct {
	BcryptManager *encrypt.BcryptManager
}

// NewServiceBag initializes and returns a ServiceBag with all dependencies wired up.
func CreateServiceBag() *ServiceBag {
	return &ServiceBag{
		BcryptManager: encrypt.CreateBcryptManager(),
	}
}