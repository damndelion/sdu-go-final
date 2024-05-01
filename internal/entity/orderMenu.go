package entity

type OrderMenu struct {
	ID      string `gorm:"primaryKey"`
	OrderID string `gorm:"not null"`
	MenuID  string `gorm:"not null"`
}
