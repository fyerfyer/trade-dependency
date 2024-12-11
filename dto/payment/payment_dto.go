package payment

type PaymentDTO struct {
	CustomerID uint64
	OrderID    uint64
	TotalPrice float32
	Status     string
	Message    string
}

type CustomerEntity struct {
	CustomerID uint64
	Balance    float32
}

type OrderEntity struct {
	OrderID    uint64
	TotalPrice float32
}

type ChargeRequest struct {
	Customer CustomerEntity
	Order    OrderEntity
}

type GetPaymentRequest struct {
	CustomerID uint64
}

type GetPaymentResponse struct {
	Payment PaymentDTO
}
