package api

import (
	"context"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/application/core/domain"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (application Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := application.db.Save(&order)

	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (application Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return domain.Order{}, nil
}
