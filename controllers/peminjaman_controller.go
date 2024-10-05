package controllers

import (
	"go-perpus/config"
	"go-perpus/models"
	"go-perpus/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePeminjaman(c *gin.Context) {
	var peminjaman models.Peminjaman
	if err := c.ShouldBindJSON(&peminjaman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	peminjaman.TanggalPinjam = time.Now()
	peminjaman.TanggalKembali = peminjaman.TanggalPinjam.AddDate(0, 0, 14)
	peminjaman.Status = "Dipinjam"
	peminjaman.Denda = 0

	tx := config.DB.Begin()
	if err := tx.Create(&peminjaman).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat peminjaman: " + err.Error()})
		return
	}

	if err := tx.Preload("Buku").Preload("User").First(&peminjaman, peminjaman.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat data peminjaman"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Peminjaman berhasil dibuat",
		"data":    peminjaman,
	})
}

func KembalikanBuku(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	peminjaman, err := services.KembalikanBuku(uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "buku sudah dikembalikan sebelumnya" {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Buku berhasil dikembalikan",
		"data":    peminjaman,
	})
}

func GetPeminjaman(c *gin.Context) {
	var peminjaman models.Peminjaman
	id := c.Param("id")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := config.DB.Preload("Buku").Preload("User").First(&peminjaman, idUint).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Peminjaman tidak ditemukan"})
		return
	}

	services.HitungDenda(&peminjaman)

	c.JSON(http.StatusOK, gin.H{"data": peminjaman})
}

func GetAllPeminjaman(c *gin.Context) {
	var peminjamans []models.Peminjaman

	if err := config.DB.Preload("Buku").Preload("User").
		Order("tanggal_pinjam desc").Find(&peminjamans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range peminjamans {
		services.HitungDenda(&peminjamans[i])
	}

	c.JSON(http.StatusOK, gin.H{"data": peminjamans})
}
