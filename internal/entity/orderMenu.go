package entity

type OrderMenu struct {
	OrderID string `gorm:"not null"`
	Order   Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MenuID  string `gorm:"not null"`
	Menu    Menu   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
