package usecase

import (
	"context"
	"fmt"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"

	"github.com/damndelion/sdu-go-final/internal/entity"
)

type OrderMenuUseCase struct {
	repo OrderMenuRepo
}

func NewOrderMenuUseCase(r OrderMenuRepo) *OrderMenuUseCase {
	return &OrderMenuUseCase{
		repo: r,
	}
}

func (mc *OrderMenuUseCase) GetOrder(ctx context.Context) ([]entity.OrderMenu, error) {
	order, err := mc.repo.GetAllOrderMenu(ctx)
	if err != nil {
		return nil, fmt.Errorf("OrderUseCase - GetOrder: %w", err)
	}

	return order, nil
}

func (mc *OrderMenuUseCase) UpdateOrderMenuItem(ctx context.Context, item dto.UpdateOrderMenuItem) (id string, err error) {
	order, err := mc.repo.UpdateOrderMenuItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("OrderUseCase - UpdateOrderItem: %w", err)
	}

	return order, nil
}

func (mc *OrderMenuUseCase) DeleteOrderItem(ctx context.Context, id string) error {
	err := mc.repo.DeleteOrderMenuItem(ctx, id)
	if err != nil {
		return fmt.Errorf("OrderUseCase - DeleteOrderItem: %w", err)
	}

	return nil
}
