package customer

import "time"

type CustomerDTO struct {
	ID        uint64
	Name      string
	Status    string
	Balance   float32
	CreatedAt time.Time
}

type CreateCustomerRequest struct {
	CustomerName string
}

type CreateCustomerResponse struct {
	CustomerID uint64
	Success    bool
}

type GetCustomerRequest struct {
	CustomerName string
}

type GetCustomerResponse struct {
	Customer *CustomerDTO
}

type DeactivateCustomer struct {
	CustomerName string
}

type DeactivateCustomerResponse struct {
	Success bool
}

type ActivateCustomer struct {
	CustomerName string
}

type ActivateCustomerResponse struct {
	Success bool
}

type OrderItem struct {
	ProductCode string
	UnitPrice   float32
	Quantity    int32
}

type SubmitOrderRequest struct {
	CustomerName string
	OrderItems   []*OrderItem
}

type SubmitOrderResponse struct {
	Success bool
	Message string
}

type PayOrderRequest struct {
	CustomerName string
	OrderID      uint64
}

type PayOrderResponse struct {
	Success bool
	Message string
}

type Order struct {
	OrderID   uint64
	Items     []*OrderItem
	Status    string
	CreatedAt time.Time
}

type GetUnpaidOrdersRequest struct {
	CustomerName string
}

type GetUnpaidOrdersResponse struct {
	UnpaidOrders []*Order
}

type StoreBalanceRequest struct {
	CustomerName string
	Balance      float32
}

type StoreBalanceResponse struct {
	Message string
}
