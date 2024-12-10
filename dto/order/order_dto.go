package order

import "time"

type OrderDTO struct {
	ID         uint64
	CustomerID uint64
	Items      []OrderItemDTO
	Status     string
	CreatedAt  time.Time
}

type OrderItemDTO struct {
	ProductCode string
	UnitPrice   float32
	Quantity    int32
}

type CustomerEntity struct {
	CustomerID uint64
	Balance    float32
}

type OrderEntity struct {
	OrderID uint64
	Items   []*OrderItemDTO
	Status  string
}

type ProcessItemsRequest struct {
	Customer   CustomerEntity
	OrderItems []*OrderItemDTO
}

type ProcessItemsResponse struct {
	Message string
}

type ProcessOrderRequest struct {
	Customer CustomerEntity
	Order    OrderEntity
}

type ProcessOrderResponse struct {
	Message string
}

type GetUnpaidOrdersRequest struct {
	CustomerID uint64
}

type GetUnpaidOrdersResponse struct {
	Orders []OrderEntity
}
