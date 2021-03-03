package data

import (
	"github.com/go-pg/pg/v10"
	"github.com/haardikk21/kruder/cmd/kruder/config"
	"github.com/sirupsen/logrus"
)

// DatabaseService represents the database class
type DatabaseService struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDatabaseService initializes a DatabaseService
func NewDatabaseService(config *config.DatabaseConfig, log *logrus.Logger) *DatabaseService {
	db := pg.Connect(&pg.Options{
		Addr:     config.DBAddress,
		User:     config.DBUsername,
		Password: config.DBPassword,
		Database: config.DBName,
	})

	log.Info("Successfully connected to database!")

	return &DatabaseService{
		log: log,
		db:  db,
	}
}
