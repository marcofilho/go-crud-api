package model

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) FindUserByID(id string) (*userDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserByID model", zap.String("journey", "findUserByID"))

	logger.Info("User found successfully", zap.String("journey", "findUserByID"))

	return &userDomain{
		email:    "",
		password: "",
		name:     "",
		age:      0,
	}, nil
}

func (ud *userDomain) FindUserByEmail(email string) (*userDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail model", zap.String("journey", "findUserByEmail"))

	logger.Info("User found successfully", zap.String("journey", "findUserByEmail"))

	return &userDomain{
		email:    "",
		password: "",
		name:     "",
		age:      0,
	}, nil
}
