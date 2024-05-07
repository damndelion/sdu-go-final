package entity

import "time"

type Order struct {
	ID          string    `gorm:"primaryKey"`
	UserID      string    `gorm:"not null"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status      string    `gorm:"not null"`
	TotalPrice  int       `gorm:"not null"`
	Timestamp   time.Time `gorm:"not null"`
	PaymentType string    `gorm:"not null"`

	MenuItems []Menu `gorm:"many2many:order_menus"`
}
