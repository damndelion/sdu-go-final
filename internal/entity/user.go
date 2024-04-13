package entity

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null";gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null";gorm:"default:user"`
}
