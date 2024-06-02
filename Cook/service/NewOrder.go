package service

import (
	"cook/db"
	"cook/types"
	"encoding/json"
)

type NewOrderService struct {
}

func (no *NewOrderService) Handle(jsonData []byte) error {
	var newOrder types.Order
	err := json.Unmarshal(jsonData, &newOrder)
	if err != nil {
		return err
	}

	for _, item := range newOrder.Items {
		newOrder.Total = newOrder.Total + item.Value
	}

	db.Orders = append(db.Orders, newOrder)

	return nil
}
