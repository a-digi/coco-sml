package config

import (
	"github.com/a-digi/coco-sml/src/db/binary"
	"github.com/a-digi/coco-sml/src/api/config/install"
)

// SetupApi initializes the API setup with a given binary.Server
func SetupApi(server *binary.Server) {
	install.SetupInternalDatabase(server)
	install.SetupUserTable(server)
	install.SetupRootUserTable(server)
}
