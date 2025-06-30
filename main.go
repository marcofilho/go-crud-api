package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcofilho/go-crud-api/src/controller/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error setting environment variable:", err)
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	router.Run(":8080")
}
