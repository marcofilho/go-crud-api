package model

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUserByID(id string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserByID model", zap.String("journey", "findUserByID"))

	logger.Info("User found successfully", zap.String("journey", "findUserByID"))

	return &UserDomain{
		Email:    "",
		Password: "",
		Name:     "",
		Age:      0,
	}, nil
}

func (ud *UserDomain) FindUserByEmail(email string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail model", zap.String("journey", "findUserByEmail"))

	logger.Info("User found successfully", zap.String("journey", "findUserByEmail"))

	return &UserDomain{
		Email:    "",
		Password: "",
		Name:     "",
		Age:      0,
	}, nil
}
