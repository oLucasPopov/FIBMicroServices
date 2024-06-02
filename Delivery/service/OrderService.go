package service

import (
	"delivery/db"
	"delivery/types"
)

type OrderService struct {
}

func (os *OrderService) NewOrder(order types.Order) (*types.Order, error) {
	for _, item := range order.Items {
		order.Total = order.Total + item.Value
	}

	db.Orders = append(db.Orders, order)

	return &order, nil
}

func (os *OrderService) ListOrders() ([]types.Order, error) {

	if db.Orders == nil {
		return make([]types.Order, 0), nil
	}

	return db.Orders, nil
}
