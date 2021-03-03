package router

import (
	"net/http"

	"github.com/haardikk21/kruder/cmd/kruder/data"
	"github.com/sirupsen/logrus"
)

// EndpointHandlerService is the REST handler for Endpoint related requests
type EndpointHandlerService struct {
	db  *data.DatabaseService
	log *logrus.Logger
}

// NewEndpointHandler returns an instance of EndpointHandlerService
func NewEndpointHandler(db *data.DatabaseService, log *logrus.Logger) *EndpointHandlerService {
	return &EndpointHandlerService{
		db:  db,
		log: log,
	}
}

// CreateEndpoint creates an endpoint
func (h *EndpointHandlerService) CreateEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))

}

// UpdateEndpoint updates an endpoint
func (h *EndpointHandlerService) UpdateEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}

// GetEndpoint gets endpoint(s)
func (h *EndpointHandlerService) GetEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}

// DeleteEndpoint deletes endpoint(s)
func (h *EndpointHandlerService) DeleteEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}
