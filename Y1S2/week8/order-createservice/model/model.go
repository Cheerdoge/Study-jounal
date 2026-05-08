package model

type Order struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Item   string
	Status string
}
