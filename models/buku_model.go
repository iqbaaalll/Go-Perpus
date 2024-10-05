package models

type Buku struct {
	ID          uint `gorm:"primaryKey"`
	Judul       string
	Pengarang   string
	Penerbit    string
	TahunTerbit int
	KategoriID  uint
	Kategori    Kategori `gorm:"foreignKey:KategoriID"`
}
