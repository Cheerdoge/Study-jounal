package model

type Order struct {
	ID     string `gorm:"primaryKey"`
	Item   string
	Status string
}
