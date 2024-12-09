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
	balance    float32
}

type OrderEntity struct {
	OrderID    uint64
	TotalPrice float32
}

type ChargeRequestDTO struct {
	Customer CustomerEntity
	Order    OrderEntity
}

type ChargeResponseDTO struct {
	Status  string
	Message string
}
