package main

import (
	"github.com/marcofilho/go-crud-api/src/controller"
	"github.com/marcofilho/go-crud-api/src/model/repository"
	"github.com/marcofilho/go-crud-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
