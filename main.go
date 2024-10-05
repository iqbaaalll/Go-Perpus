package main

import (
	"go-perpus/config"
	"go-perpus/controllers"
	"go-perpus/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Kategori{}, &models.Buku{}, &models.Peminjaman{})

	router := gin.Default()

	// User Routes
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/user/:id", controllers.GetUserById)
	router.POST("/delete-user/:id", controllers.DeleteUser)
	router.POST("/update-user/:id", controllers.UpdateUser)

	// Kategori Routes
	router.POST("/kategori", controllers.CreateKategori)
	router.GET("/kategori", controllers.GetAllKategori)
	router.GET("/kategori/:id", controllers.GetKategoriById)

	//Buku Routes
	router.POST("/buku", controllers.CreateBuku)
	router.GET("/buku", controllers.GetAllBuku)
	router.GET("/buku/:id", controllers.GetBukuById)
	router.POST("/delete-buku/:id", controllers.DeleteBuku)
	router.POST("/update-buku/:id", controllers.UpdateBuku)

	//Peminjaman Routes
	router.POST("/peminjaman", controllers.CreatePeminjaman)
	router.PUT("/peminjaman/:id/kembali", controllers.KembalikanBuku)
	router.GET("/peminjaman/:id", controllers.GetPeminjaman)
	router.GET("/peminjaman", controllers.GetAllPeminjaman)

	router.Run(":3000")
}
