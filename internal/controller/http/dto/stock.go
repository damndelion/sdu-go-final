package dto

type CreateStockItem struct {
	MenuID   string `json:"menu_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdateStockItem struct {
	ID       string `json:"id" binding:"required"`
	MenuID   string `json:"menu_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}
