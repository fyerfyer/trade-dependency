package order

import "time"

type OrderDTO struct {
	ID         uint64         `json:"id"`
	CustomerID uint64         `json:"customer_id"`
	Items      []OrderItemDTO `json:"items"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
}

type OrderItemDTO struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type ProcessItemsRequest struct {
	CustomerId uint64
	OrderItems []*OrderItemDTO
}

type ProcessItemsResponse struct {
	Message string
}

type ProcessOrderRequest struct {
	Order OrderDTO
}

type ProcessOrderResponse struct {
	Message string
}
