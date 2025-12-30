package route

import (
	"net/http"
)

// RouteBuilder helps to register API routes and handlers

type RouteBuilder struct {
	mux *http.ServeMux
}

// NewRouteBuilder creates a new RouteBuilder with a fresh ServeMux
func NewRouteBuilder() *RouteBuilder {
	return &RouteBuilder{
		mux: http.NewServeMux(),
	}
}

// Handle registers a handler for a given pattern
func (rb *RouteBuilder) Handle(pattern string, handler http.HandlerFunc) {
	rb.mux.HandleFunc(pattern, handler)
}

// Handler returns the underlying ServeMux for use in http.Server
func (rb *RouteBuilder) Handler() http.Handler {
	return rb.mux
}