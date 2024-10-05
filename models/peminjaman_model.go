package models

import "time"

type Peminjaman struct {
	ID                  uint `gorm:"primaryKey"`
	BukuID              uint
	Buku                Buku `gorm:"foreignKey:BukuID"`
	UserID              uint
	User                User `gorm:"foreignKey:UserID"`
	TanggalPinjam       time.Time
	TanggalKembali      time.Time
	TanggalDikembalikan *time.Time
	Denda               float64
	Status              string
}
