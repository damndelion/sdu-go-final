package usecase

import (
	"context"
	"fmt"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"

	"github.com/damndelion/sdu-go-final/internal/entity"
)

type StockUseCase struct {
	repo StockRepo
}

func NewStockUseCase(r StockRepo) *StockUseCase {
	return &StockUseCase{
		repo: r,
	}
}

func (mc *StockUseCase) GetStock(ctx context.Context) ([]entity.Stock, error) {
	stock, err := mc.repo.GetAllStock(ctx)
	if err != nil {
		return nil, fmt.Errorf("StockUseCase - GetStock: %w", err)
	}

	return stock, nil
}

func (mc *StockUseCase) CreateStockItem(ctx context.Context, item dto.CreateStockItem) (id string, err error) {
	stock, err := mc.repo.CreateStockItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("StockUseCase - CreateStockItem: %w", err)
	}

	return stock, nil
}

func (mc *StockUseCase) UpdateStockItem(ctx context.Context, item dto.UpdateStockItem) (id string, err error) {
	stock, err := mc.repo.UpdateStockItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("StockUseCase - UpdateStockItem: %w", err)
	}

	return stock, nil
}

func (mc *StockUseCase) DeleteStockItem(ctx context.Context, id string) error {
	err := mc.repo.DeleteStockItem(ctx, id)
	if err != nil {
		return fmt.Errorf("StockUseCase - DeleteStockItem: %w", err)
	}

	return nil
}
