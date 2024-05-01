package repo

import "C"
import (
	"context"
	"errors"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepo struct {
	db *gorm.DB
}

func NewStockRepo(pg *gorm.DB) *StockRepo {
	return &StockRepo{pg}
}

func (r *StockRepo) GetAllStock(ctx context.Context) (stock []entity.Stock, err error) {
	res := r.db.Find(&stock)
	if res.Error != nil {
		return nil, res.Error
	}

	return stock, nil
}

func (r *StockRepo) CreateStockItem(ctx context.Context, items dto.CreateStockItem) (id string, err error) {
	stockUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	stock := entity.Stock{
		ID:       stockUuid.String(),
		MenuID:   items.MenuID,
		Quantity: items.Quantity,
	}

	res := r.db.Create(&stock)
	if res.Error != nil {
		return "", res.Error
	}

	return stock.ID, nil
}

func (r *StockRepo) UpdateStockItem(ctx context.Context, item dto.UpdateStockItem) (id string, err error) {
	stock := entity.Stock{
		MenuID:   item.MenuID,
		Quantity: item.Quantity,
	}
	res := r.db.Model(&stock).Where("id = ?", item.ID).Updates(&stock)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			return "", errors.New("stock with that name already exists")
		}
		return "", res.Error
	}

	return item.ID, nil
}

func (r *StockRepo) DeleteStockItem(ctx context.Context, id string) error {
	res := r.db.Model(&entity.Stock{}).Where("id = ?", id).Delete(&entity.Stock{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
