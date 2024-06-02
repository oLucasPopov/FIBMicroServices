package types

type Order struct {
	OrderId int64   `json:"orderId"`
	Items   []Item  `json:"items"`
	Total   float32 `json:"total"`
}
