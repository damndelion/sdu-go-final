package dto

import (
	"github.com/damndelion/sdu-go-final/internal/entity"
	"time"
)

type CreateOrderItemRequest struct {
	Items       []CreateOrderItem `json:"items" binding:"required"`
	PaymentType string            `json:"payment_type" binding:"required"`
}

type CreateOrderItem struct {
	MenuID string `json:"menu_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Price  int    `json:"price" binding:"required"`
}

type UpdateOrderItem struct {
	ID          string    `json:"id" binding:"required"`
	UserID      string    `json:"user_id" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	TotalPrice  int       `json:"total_price" binding:"required"`
	Timestamp   time.Time `json:"timestamp" binding:"required"`
	PaymentType string    `json:"payment_type" binding:"required"`
}

type OrderWithMenuInfo struct {
	entity.Order                // Embed the entire Order entity
	MenuItems    []MenuItemInfo // Slice to hold MenuItem details
}
type MenuItemInfo struct {
	MenuID int    `json:"menu_id"`
	Name   string `json:"name"`
	// Add other relevant menu information fields here
}
