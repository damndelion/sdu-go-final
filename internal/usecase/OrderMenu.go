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

func (mc *OrderMenuUseCase) GetOrderMenu(ctx context.Context) ([]entity.OrderMenu, error) {
	orderMenu, err := mc.repo.GetAllOrderMenu(ctx)
	if err != nil {
		return nil, fmt.Errorf("OrderUseCase - GetOrder: %w", err)
	}

	return orderMenu, nil
}

func (mc *OrderMenuUseCase) UpdateOrderMenuItem(ctx context.Context, item dto.UpdateOrderMenuItem) (id string, err error) {
	orderMenu, err := mc.repo.UpdateOrderMenuItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("OrderMenuUseCase - UpdateOrderMenuItem: %w", err)
	}

	return orderMenu, nil
}

func (mc *OrderMenuUseCase) DeleteOrderMenuItem(ctx context.Context, id string) error {
	err := mc.repo.DeleteOrderMenuItem(ctx, id)
	if err != nil {
		return fmt.Errorf("OrderMenuUseCase - DeleteOrderMenuItem: %w", err)
	}

	return nil
}
