package dto

import (
	"github.com/damndelion/sdu-go-final/internal/entity"
	"time"
)

type CreateOderItem struct {
	UserID      entity.User `json:"user_id" binding:"required"`
	Status      string      `json:"status" binding:"required"`
	TotalPrice  int         `json:"total_price" binding:"required"`
	Timestamp   time.Time   `json:"timestamp" binding:"required"`
	PaymentType string      `json:"payment_type" binding:"required"`
}

type UpdateOrderItem struct {
	ID          string      `json:"id" binding:"required"`
	UserID      entity.User `json:"user_id" binding:"required"`
	Status      string      `json:"status" binding:"required"`
	TotalPrice  int         `json:"total_price" binding:"required"`
	Timestamp   time.Time   `json:"timestamp" binding:"required"`
	PaymentType string      `json:"payment_type" binding:"required"`
}
