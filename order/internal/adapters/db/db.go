package db

import (
	"fmt"
	"github.com/stand-sure/grpc-microservices-in-go/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerId int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderId     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})

	if openErr != nil {
		return nil, fmt.Errorf("db migration error: %v", openErr)
	}

	err := db.AutoMigrate(&Order{}, OrderItem{})

	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntry Order

	result := a.db.First(&orderEntry, id)

	var orderItems []domain.OrderItem

	for _, orderItem := range orderEntry.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntry.ID),
		CustomerID: orderEntry.CustomerId,
		Status:     orderEntry.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntry.CreatedAt.UnixNano(),
	}

	return order, result.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem

	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	orderModel := Order{
		CustomerId: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	result := a.db.Create(&orderModel)
	if result.Error == nil {
		order.ID = int64(orderModel.ID)
	}

	return result.Error
}
