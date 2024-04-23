package dto

type CreateMenuItem struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required,email"`
}

type UpdateMenuItem struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required,email"`
}
