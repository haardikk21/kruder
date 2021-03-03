package router

import (
	"net/http"

	"github.com/haardikk21/kruder/cmd/kruder/data"
	"github.com/sirupsen/logrus"
)

// SchemaHandlerService is the REST handler for Schema related requests
type SchemaHandlerService struct {
	db  *data.DatabaseService
	log *logrus.Logger
}

// NewSchemaHandler returns an instance of SchemaHandler
func NewSchemaHandler(db *data.DatabaseService, log *logrus.Logger) *SchemaHandlerService {
	return &SchemaHandlerService{
		db:  db,
		log: log,
	}
}

// CreateSchema creates a schema
func (h *SchemaHandlerService) CreateSchema(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}

// UpdateSchema updates a schema
func (h *SchemaHandlerService) UpdateSchema(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}

// GetSchema returns schema(s)
func (h *SchemaHandlerService) GetSchema(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}

// DeleteSchema deletes schema(s)
func (h *SchemaHandlerService) DeleteSchema(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}
