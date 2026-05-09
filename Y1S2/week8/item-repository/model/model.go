package model

type Item struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Num  int    `gorm:"not null"`
}

type ProcessedOrder struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"not null"`
}
