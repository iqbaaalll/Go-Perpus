package controllers

import (
	"go-perpus/config"
	"go-perpus/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateKategori(c *gin.Context) {
	var kategori models.Kategori
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&kategori).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori created successfully!"})
}

func GetAllKategori(c *gin.Context) {
	var kategoris []models.Kategori
	config.DB.Find(&kategoris)
	c.JSON(http.StatusOK, kategoris)
}

func GetKategoriById(c *gin.Context) {
	var kategori models.Kategori

	id := c.Param("id")

	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found"})
		return
	}

	c.JSON(http.StatusOK, kategori)
}
