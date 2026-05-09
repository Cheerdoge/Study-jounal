package model

type Order struct {
	ID       uint `gorm:"primaryKey"`
	ItemName string
	Nums     int
}
