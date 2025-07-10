package service

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "deleteUser"))

	logger.Info("User deleted successfully", zap.String("userID", id))

	return nil
}
