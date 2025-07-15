package service

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser service",
		zap.String("journey", "updateUser"))

	domain, err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser repository", err,
			zap.String("journey", "updateUser"))
		return nil, err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("userID", id),
		zap.String("journey", "updateUser"),
	)

	return nil
}
