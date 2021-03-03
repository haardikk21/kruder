package main

import (
	"github.com/haardikk21/kruder/cmd/kruder/config"
	"github.com/haardikk21/kruder/cmd/kruder/data"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.Logger{}
	log.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
	}

	// Initialize the configuration
	conf, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Error("Could not load environment variables")
	}

	// Initialize the database service
	dbService := data.NewDatabaseService(&conf.DatabaseConfig, &log)

	// Initialize the router
	routerService := NewRouterService(&conf.RouterConfig, dbService, &log)

	routerService.Run()
}
