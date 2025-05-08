// Package order implements the order processing and management functionality
package order

import (
	"errors"

	"github.com/Vractos/ecoffe-go/entity"
	"github.com/Vractos/ecoffe-go/usecases/common"
)

// Error definitions for order-related operations
var (
	// ErrOrderNotFound is returned when an order cannot be found
	ErrOrderNotFound = errors.New("order not found")
)

type OrderService struct {
	repo   Repository    // Order repository for data persistence
	logger common.Logger // Logger for error and info logging
}

func NewOrderService(
	repository Repository,
	logger common.Logger,
) *OrderService {
	return &OrderService{
		repo:   repository,
		logger: logger,
	}
}

func (o *OrderService) CreateOrder(input CreateOrderDtoInput) (*entity.ID, error) {
	order, err := entity.NewOrder(
		input.Client,
		input.Item,
		input.Quantity,
		input.Observation,
		entity.Pending,
	)
	if err != nil {
		o.logger.Error("Failed to create order", err)
		return &order.ID, err
	}

	err = o.repo.CreateOrder(order)
	if err != nil {
		o.logger.Error("Failed to save order", err)
		return &order.ID, err
	}
	return &order.ID, nil
}

func (o *OrderService) RetrieveAllOrders() (*[]entity.Order, error) {
	orders, err := o.repo.RetrieveAllOrders()
	if err != nil {
		o.logger.Error("Failed to retrieve orders", err)
		return nil, err
	}
	return orders, nil
}
