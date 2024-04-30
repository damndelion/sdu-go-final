package repo

import "C"
import (
	"context"
	"errors"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"

	"github.com/damndelion/sdu-go-final/internal/entity"
)

type OrderRepo struct {
	db *gorm.DB
}

// NewOderRepo
func NewOderRepo(pg *gorm.DB) *OrderRepo {
	return &OrderRepo{pg}
}

func (r *OrderRepo) GetAllOrder(ctx context.Context) (order []entity.Order, err error) {
	res := r.db.Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	return order, nil
}

func (r *OrderRepo) CreateOrderItem(ctx context.Context, item dto.CreateOrderItemRequest, userID string) (id string, err error) {
	orderUuid, err := uuid.NewRandom()
	var orderItemData dto.CreateOrderItemRequest
	totalPrice := 0
	for _, item := range orderItemData.Items {
		totalPrice += item.Price
	}
	if err != nil {
		return "", err
	}
	order := entity.Order{
		ID:          orderUuid.String(),
		UserID:      userID,
		Status:      "crated",
		TotalPrice:  totalPrice,
		Timestamp:   time.Now(),
		PaymentType: item.PaymentType,
	}
	res := r.db.Create(&order)
	if res.Error != nil {
		return "", res.Error
	}

	return order.ID, nil
}

func (r *OrderRepo) UpdateOrderItem(ctx context.Context, item dto.UpdateOrderItem) (id string, err error) {
	order := entity.Order{
		UserID:      item.UserID,
		Status:      item.Status,
		TotalPrice:  item.TotalPrice,
		Timestamp:   item.Timestamp,
		PaymentType: item.PaymentType,
	}
	res := r.db.Model(&order).Where("id = ?", item.ID).Updates(&order)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return "", errors.New("order with that name already exists")
		}
		return "", res.Error
	}

	return item.Status, nil
}

func (r *OrderRepo) DeleteOrderItem(ctx context.Context, id string) error {
	res := r.db.Model(&entity.Order{}).Where("id = ?", id).Delete(&entity.Order{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
