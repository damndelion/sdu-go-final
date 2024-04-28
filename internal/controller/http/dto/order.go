package dto

type CreateOrderItemRequest struct {
	Items       []CreateMenuItem `json:"items" binding:"required"`
	PaymentType string           `json:"payment_type" binding:"required"`
}

type CreateOrderItem struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
