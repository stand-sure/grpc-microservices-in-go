package grpc

import (
	"fmt"
	"github.com/huseyinbabal/microservices-proto/golang/order"
	log "github.com/sirupsen/logrus"
	"github.com/stand-sure/grpc-microservices-in-go/order/config"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	order.RegisterOrderServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to start grpc server, error: %v", err)
	}
}
