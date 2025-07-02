package model

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) UpdateUser(id string) *rest_err.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "updateUser"))

	ud.EncryptPassword()

	logger.Info("User updated successfully", zap.String("journey", "updateUser"))

	return nil
}
