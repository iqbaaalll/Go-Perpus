package controllers

import (
	"go-perpus/config"
	"go-perpus/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBuku(c *gin.Context) {
	var buku models.Buku
	if err := c.ShouldBindJSON(&buku); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku created successfully!"})
}

func GetAllBuku(c *gin.Context) {
	var bukus []models.Buku
	config.DB.Find(&bukus)
	c.JSON(http.StatusOK, bukus)
}

func GetBukuById(c *gin.Context) {
	var buku models.Buku
	id := c.Param("id")

	if err := config.DB.First(&buku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku not found"})
		return
	}

	c.JSON(http.StatusOK, buku)
}

func DeleteBuku(c *gin.Context) {
	var buku models.Buku
	id := c.Param("id")

	if err := config.DB.First(&buku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku not found"})
		return
	}

	if err := config.DB.Delete(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku deleted succsessfully"})
}

func UpdateBuku(c *gin.Context) {
	id := c.Param("id")

	var existingBuku models.Buku
	if err := config.DB.First(&existingBuku, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku not found"})
		return
	}

	var updateData models.Buku
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&existingBuku).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "buku": existingBuku})

}
