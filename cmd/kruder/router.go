package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haardikk21/kruder/cmd/kruder/config"
	"github.com/haardikk21/kruder/cmd/kruder/data"
	"github.com/haardikk21/kruder/cmd/kruder/router"
	"github.com/sirupsen/logrus"
)

// RouterService represents the router class
type RouterService struct {
	router *mux.Router
	config *config.RouterConfig
	db     *data.DatabaseService
	log    *logrus.Logger
}

// NewRouterService returns an instance of RouterService
func NewRouterService(config *config.RouterConfig, db *data.DatabaseService, log *logrus.Logger) *RouterService {
	r := mux.NewRouter()

	// Initialize schema routes
	schemaHandler := router.NewSchemaHandler(db, log)
	r.HandleFunc("/v1/schema", schemaHandler.CreateSchema).Methods("POST")
	r.HandleFunc("/v1/schema", schemaHandler.GetSchema).Methods("GET")
	r.HandleFunc("/v1/schema", schemaHandler.UpdateSchema).Methods("PUT")
	r.HandleFunc("/v1/schema", schemaHandler.DeleteSchema).Methods("DELETE")

	// Initialize endpoint routes
	endpointHandler := router.NewEndpointHandler(db, log)
	r.HandleFunc("/v1/endpoint", endpointHandler.CreateEndpoint).Methods("POST")
	r.HandleFunc("/v1/endpoint", endpointHandler.GetEndpoint).Methods("GET")
	r.HandleFunc("/v1/endpoint", endpointHandler.UpdateEndpoint).Methods("PUT")
	r.HandleFunc("/v1/endpoint", endpointHandler.DeleteEndpoint).Methods("DELETE")

	return &RouterService{
		router: r,
		config: config,
		db:     db,
		log:    log,
	}
}

// Run starts the HTTP server
func (r *RouterService) Run() {
	r.log.Fatal(http.ListenAndServe(fmt.Sprint(r.config.Port), r.router))
}
