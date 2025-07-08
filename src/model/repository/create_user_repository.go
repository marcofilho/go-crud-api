package repository

import (
	"context"
	"os"

	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
)

var (
	COLLECTION_NAME = "COLLECTION_NAME"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")
	ctx := context.Background()

	collection_name := os.Getenv("COLLECTION_NAME")
	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValeu()
	if err != nil {
		logger.Error("Error getting JSON value from user domain:", err)
		return nil, rest_err.NewInternalServerError("Error getting JSON value from user domain: " + err.Error())
	}

	result, err := collection.InsertOne(ctx, value)
	if err != nil {
		logger.Error("Error inserting user into database:", err)
		return nil, rest_err.NewInternalServerError("Error inserting user into database: " + err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
