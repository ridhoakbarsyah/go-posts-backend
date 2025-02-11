package main

import (
	"fmt"
	"go-posts-backend/config"
	"go-posts-backend/models"
)

func main() {
	// Koneksi ke database
	config.ConnectDatabase()

	// Jalankan migrasi database
	err := config.DB.AutoMigrate(&models.Post{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi:", err)
	} else {
		fmt.Println("Migrasi berhasil!")
	}
}
