package status

import (
	"net/http"
	"github.com/a-digi/coco-sml/src/api/config/di"
	"github.com/a-digi/coco-sml/src/api/response"
)

// StatusHandler handles GET /v1/status requests
func StatusHandler(w http.ResponseWriter, r *http.Request, serviceBag *di.ServiceBag) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response.ErrorResponse("Method not allowed"))
		return
	}

	w.Write(response.SuccessResponse(map[string]string{"status": "ok", "server": "coco-sml"}))
}
