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

func (r *OrderRepo) CreateOrderItem(ctx context.Context, items dto.CreateOrderItemRequest, userID string) (id string, err error) {
	orderUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	totalPrice := 0
	for _, item := range items.Items {
		totalPrice += item.Price
	}
	order := entity.Order{
		ID:          orderUuid.String(),
		UserID:      userID,
		Status:      "crated",
		TotalPrice:  totalPrice,
		Timestamp:   time.Now(),
		PaymentType: items.PaymentType,
	}

	for _, item := range items.Items {
		orderMenuUuid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		orderMenu := entity.OrderMenu{
			ID:      orderMenuUuid.String(),
			OrderID: orderUuid.String(),
			MenuID:  item.MenuID,
		}
		create := r.db.Create(&orderMenu)
		if create.Error != nil {
			return "", create.Error
		}
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
