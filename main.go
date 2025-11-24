package main

import (
	"log"
	"os"
	"stocky-backend/routes"
	"stocky-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	utils.InitLogger()
	utils.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	utils.Logger.Infof("Server started on port %s", port)
	r.Run(":" + port)
}
