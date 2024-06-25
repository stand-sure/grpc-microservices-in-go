package ports

import (
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/application/core/domain"
)

type DBPort interface {
	Save(order *domain.Order) error
	Get(id string) (domain.Order, error)
}
