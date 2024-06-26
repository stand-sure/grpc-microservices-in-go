package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stand-sure/grpc-microservices-in-go/order/config"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/adapters/db"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/adapters/grpc"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
