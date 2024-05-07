package usecase

import (
	"context"
	"fmt"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"

	"github.com/damndelion/sdu-go-final/internal/entity"
)

type OrderUseCase struct {
	repo OderRepo
}

func NewOrderUseCase(r OderRepo) *OrderUseCase {
	return &OrderUseCase{
		repo: r,
	}
}

func (mc *OrderUseCase) GetOrder(ctx context.Context) ([]entity.Order, error) {
	order, err := mc.repo.GetAllOrder(ctx)
	if err != nil {
		return nil, fmt.Errorf("OrderUseCase - GetOrder: %w", err)
	}

	return order, nil
}

func (mc *OrderUseCase) GetUserCurrentOrder(ctx context.Context, userId string) ([]entity.Order, error) {
	order, err := mc.repo.GetUserCurrentOrders(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("OrderUseCase - GetUserCurrentOrder: %w", err)
	}

	return order, nil
}
func (mc *OrderUseCase) GetUserAllOrder(ctx context.Context, userId string) ([]entity.Order, error) {
	order, err := mc.repo.GetUserCurrentOrders(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("OrderUseCase - GetUserAllOrder: %w", err)
	}

	return order, nil
}

func (mc *OrderUseCase) CreateOrderItem(ctx context.Context, item dto.CreateOrderItemRequest, userId string) (id string, err error) {
	order, err := mc.repo.CreateOrderItem(ctx, item, userId)
	if err != nil {
		return "", fmt.Errorf("OrderUseCase - CreateOrderItem: %w", err)
	}

	return order, nil
}

func (mc *OrderUseCase) UpdateOrderItem(ctx context.Context, item dto.UpdateOrderItem) (id string, err error) {
	order, err := mc.repo.UpdateOrderItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("OrderUseCase - UpdateOrderItem: %w", err)
	}

	return order, nil
}

func (mc *OrderUseCase) DeleteOrderItem(ctx context.Context, id string) error {
	err := mc.repo.DeleteOrderItem(ctx, id)
	if err != nil {
		return fmt.Errorf("OrderUseCase - DeleteOrderItem: %w", err)
	}

	return nil
}
