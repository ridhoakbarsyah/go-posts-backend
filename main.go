package main

import (
	"go-posts-backend/config"
	"go-posts-backend/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}
