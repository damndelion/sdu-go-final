package entity

import "time"

type Order struct {
	ID          string    `gorm:"primaryKey"`
	UserID      string    `gorm:"not null"`
	Status      string    `gorm:"not null"`
	TotalPrice  int       `gorm:"not null"`
	Timestamp   time.Time `gorm:"not null"`
	PaymentType string    `gorm:"not null"`

	Menu []*Menu `gorm:"many2many:order_menus"`
}
