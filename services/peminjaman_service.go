package services

import (
	"errors"
	"go-perpus/config"
	"go-perpus/models"
	"time"
)

const DENDA_PER_HARI = 5000

func HitungDenda(peminjaman *models.Peminjaman) float64 {
	var denda float64

	if peminjaman.TanggalDikembalikan == nil {
		today := time.Now()
		if today.After(peminjaman.TanggalKembali) {
			hari := today.Sub(peminjaman.TanggalKembali).Hours() / 24
			denda = float64(int(hari)) * DENDA_PER_HARI
		}
	} else {
		if peminjaman.TanggalDikembalikan.After(peminjaman.TanggalKembali) {
			hari := peminjaman.TanggalDikembalikan.Sub(peminjaman.TanggalKembali).Hours() / 24
			denda = float64(int(hari)) * DENDA_PER_HARI
		}
	}

	peminjaman.Denda = denda
	return denda
}

func KembalikanBuku(peminjamanID uint) (*models.Peminjaman, error) {
	var peminjaman models.Peminjaman

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.First(&peminjaman, peminjamanID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if peminjaman.Status == "Dikembalikan" {
		tx.Rollback()
		return nil, errors.New("buku sudah dikembalikan sebelumnya")
	}

	now := time.Now()
	peminjaman.TanggalDikembalikan = &now
	peminjaman.Status = "Dikembalikan"
	HitungDenda(&peminjaman)

	if err := tx.Save(&peminjaman).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &peminjaman, nil
}
