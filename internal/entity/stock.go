package entity

type Stock struct {
	ID       string `gorm:"primaryKey"`
	MenuID   string `gorm:"not null"`
	Quantity int    `gorm:"not null"`
}
