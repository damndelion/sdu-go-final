package entity

type Menu struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"`
	Price       int    `gorm:"not null"`
	Description string `gorm:"not null"`

	Order []*Order `gorm:"many2many:order_menus"`
}
