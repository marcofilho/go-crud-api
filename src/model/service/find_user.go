package service

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByID(id string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID model", zap.String("journey", "findUserByID"))

	logger.Info("User found successfully", zap.String("journey", "findUserByID"))

	return nil, nil
}

func (ud *userDomainService) FindUserByEmail(email string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail model", zap.String("journey", "findUserByEmail"))

	logger.Info("User found successfully", zap.String("journey", "findUserByEmail"))

	return nil, nil
}
