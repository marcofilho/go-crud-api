package repository

import (
	"context"
	"os"

	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
	"github.com/marcofilho/go-crud-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"))

	ctx := context.Background()
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	filter := bson.M{"_id": id}

	_, err := collection.UpdateOne(ctx, filter, value)
	if err != nil {
		logger.Error("Error trying to update a new user", err,
			zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateUser repository executed successfully",
		zap.String("userID", value.ID.Hex()),
		zap.String("journey", "updateUser"),
	)

	return nil
}
