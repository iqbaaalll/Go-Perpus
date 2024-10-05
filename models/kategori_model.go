package models

type Kategori struct {
	ID   uint   `gorm:"primaryKey"`
	Nama string `gorm:"unique"`
}
