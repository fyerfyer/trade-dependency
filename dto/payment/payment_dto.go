package payment

type PaymentDTO struct {
	CustomerID uint64  `json:"customer_id"`
	OrderID    uint64  `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	Status     string  `json:"status"`
	Message    string  `json:"message"`
}

type ChargeRequestDTO struct {
	CustomerID uint64
	OrderId    uint64
	TotalPrice float32
}

type ChargeResponseDTO struct {
	Status  string
	Message string
}
