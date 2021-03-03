package router

import (
	"net/http"

	"github.com/haardikk21/kruder/cmd/kruder/data"
	"github.com/sirupsen/logrus"
)

// AuthHandlerService is the REST handler for auth related endpoints
type AuthHandlerService struct {
	db  *data.DatabaseService
	log *logrus.Logger
}

// NewAuthHandlerService returns an instance of AuthHandlerService
func NewAuthHandlerService(db *data.DatabaseService, log *logrus.Logger) *AuthHandlerService {
	return &AuthHandlerService{
		db:  db,
		log: log,
	}
}

// Signup is to signup for Kruder
func (h *AuthHandlerService) Signup(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}

// Login is to login to Kruder
func (h *AuthHandlerService) Login(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Not implemented"))
}
