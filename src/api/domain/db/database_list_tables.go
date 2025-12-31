package db

import (
	"log"
	"net/http"
	"github.com/a-digi/coco-sml/src/api/config/di"
	"github.com/a-digi/coco-sml/src/api/response"
	"github.com/a-digi/coco-sml/src/db/binary/orm"
)

// DatabaseListTablesHandler handles GET requests to list all tables in a database.
func DatabaseListTablesHandler(w http.ResponseWriter, r *http.Request, serviceBag *di.ServiceBag) {
	log.Println("[DatabaseListTablesHandler] called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		log.Println("[DatabaseListTablesHandler] Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response.ErrorResponse("Method not allowed"))
		return
	}

	dbName := r.URL.Query().Get("database")
	if dbName == "" {
		log.Println("[DatabaseListTablesHandler] Missing 'database' query parameter")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response.ErrorResponse("Missing 'database' query parameter"))
		return
	}

	rm, ok := serviceBag.RegistryManager.(*orm.RegistryManager)
	if !ok || rm == nil {
		log.Println("[DatabaseListTablesHandler] RegistryManager not available")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response.ErrorResponse("RegistryManager not available"))
		return
	}

	log.Printf("[DatabaseListTablesHandler] Getting EntityManager for db: %s\n", dbName)
	entityManager, err := rm.GetManager(dbName)
	if err != nil {
		log.Printf("[DatabaseListTablesHandler] Database not found: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write(response.ErrorResponse("Database not found: " + err.Error()))
		return
	}

	log.Println("[DatabaseListTablesHandler] Calling ListTables on EntityManager")
	tables, err := entityManager.ListTables()
	if err != nil {
		log.Printf("[DatabaseListTablesHandler] Failed to list tables: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response.ErrorResponse("Failed to list tables: " + err.Error()))
		return
	}

	log.Printf("[DatabaseListTablesHandler] Returning %d tables\n", len(tables))

	w.WriteHeader(http.StatusOK)
	w.Write(response.SuccessResponse(map[string]interface{}{
		"tables": tables,
	}))
}
