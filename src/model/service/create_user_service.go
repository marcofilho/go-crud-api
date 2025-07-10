package service

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser service",
		zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call CreateUser repository", err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userID", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}
