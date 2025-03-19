package model

type Book struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Title     string  `gorm:"type:varchar(120);not null"`
	Author    string  `gorm:"type:varchar(120);not null"`
	Publisher string  `gorm:"type:varchar(120);not null"`
	Country   string  `gorm:"type:varchar(120);not null"`
	Price     float64 `gorm:"type:decimal(10,2)"`
	Currency  string  `gorm:"type:varchar(10);not null"`
}
