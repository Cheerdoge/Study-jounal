package model

type Order struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	ItemName string
	Nums     int
}
