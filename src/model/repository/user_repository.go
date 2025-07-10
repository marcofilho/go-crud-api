package repository

import (
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	COLLECTION_NAME = "COLLECTION_NAME"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr)
}
