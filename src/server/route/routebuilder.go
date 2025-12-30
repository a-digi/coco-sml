package route

import (
	"net/http"

	"github.com/a-digi/coco-sml/src/db/binary"
)

// RouteBuilder helps to register API routes and handlers

type RouteBuilder struct {
	mux      *http.ServeMux
	DbServer *binary.Server
}

// NewRouteBuilder creates a new RouteBuilder with a fresh ServeMux and expects a Server argument
func NewRouteBuilder(server *binary.Server) *RouteBuilder {

	return &RouteBuilder{
		mux:      http.NewServeMux(),
		DbServer: server,
	}
}

// ServerHandlerFunc is a handler function that has access to the binary.Server
type ServerHandlerFunc func(w http.ResponseWriter, r *http.Request, srv *binary.Server)

// Handle registers a handler for a given pattern, providing DbServer as additional argument
func (rb *RouteBuilder) Handle(pattern string, handler ServerHandlerFunc) {
	rb.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, rb.DbServer)
	})
}

// Handler returns the underlying ServeMux for use in http.Server
func (rb *RouteBuilder) Handler() http.Handler {
	return rb.mux
}