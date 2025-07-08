package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcofilho/go-crud-api/src/configuration/database/mongodb"
	"github.com/marcofilho/go-crud-api/src/controller"
	"github.com/marcofilho/go-crud-api/src/controller/routes"
	"github.com/marcofilho/go-crud-api/src/model/service"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error setting environment variable:", err)
	}

	mongodb.NewMongoDBConnection(context.Background())

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
