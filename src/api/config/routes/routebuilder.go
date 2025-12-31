package routes

import (
	"net/http"

	"github.com/a-digi/coco-sml/src/api/config/di"
)

// RouteBuilder helps to register API routes and handlers

type RouteBuilder struct {
	mux        *http.ServeMux
	ServiceBag *di.ServiceBag
}

// NewRouteBuilder creates a new RouteBuilder with a fresh ServeMux and expects a ServiceBag argument
func NewRouteBuilder(serviceBag *di.ServiceBag) *RouteBuilder {

	return &RouteBuilder{
		mux:        http.NewServeMux(),
		ServiceBag: serviceBag,
	}
}

// ServiceBagHandlerFunc is a handler function that has access to the ServiceBag
type ServiceBagHandlerFunc func(w http.ResponseWriter, r *http.Request, sb *di.ServiceBag)

// Handle registers a handler for a given pattern, providing ServiceBag as additional argument
func (rb *RouteBuilder) Handle(pattern string, handler ServiceBagHandlerFunc) {
	rb.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, rb.ServiceBag)
	})
}

// Handler returns the underlying ServeMux for use in http.Server
func (rb *RouteBuilder) Handler() http.Handler {
	return rb.mux
}