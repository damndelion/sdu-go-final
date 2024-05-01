package dto

type UpdateOrderMenuItem struct {
	ID      string `json:"id" binding:"required"`
	OrderID string `json:"order-id" binding:"required"`
	MenuID  string `json:"menu_id" binding:"required"`
}
