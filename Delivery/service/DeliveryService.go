package service

import "delivery/types"

type DeliveryService struct {
}

func (ds *DeliveryService) ListMenu() ([]types.Item, error) {
	return []types.Item{
		{
			Id:    1,
			Name:  "Coxinha",
			Value: 7.00,
		},
		{
			Id:    2,
			Name:  "Pizza Augusta",
			Value: 80.00,
		},
		{
			Id:    3,
			Name:  "Pastel",
			Value: 11.90,
		},
		{
			Id:    4,
			Name:  "Lasanha",
			Value: 27.00,
		},
	}, nil
}
