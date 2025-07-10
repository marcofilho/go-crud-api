package service

import (
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID service",
		zap.String("journey", "findUserByID"))

	userDomain, err := ud.userRepository.FindUserByID(id)
	if err != nil {
		logger.Error("Error trying to call FindByUserID repository", err,
			zap.String("userID", id),
			zap.String("journey", "findUserByID"))
		return nil, err
	}

	logger.Info("FindByUserID service executed successfully",
		zap.String("userID", userDomain.GetID()),
		zap.String("journey", "createUser"))

	return userDomain, nil
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service",
		zap.String("journey", "findUserByEmail"))

	userDomain, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail repository", err,
			zap.String("email", email),
			zap.String("journey", "findUserByID"))
		return nil, err
	}

	logger.Info("FindUserByEmail service executed successfully",
		zap.String("email", email),
		zap.String("userID", userDomain.GetID()),
		zap.String("journey", "createUser"))

	return userDomain, nil
}

func (ud *userDomainService) FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllUsers service",
		zap.String("journey", "findAllUsers"))

	usersDomain, err := ud.userRepository.FindAllUsers()
	if err != nil {
		logger.Error("Error trying to call FindAllUsers repository", err,
			zap.String("journey", "findAllUsers"))
		return nil, err
	}

	logger.Info("FindAllUsers service executed successfully",
		zap.Int("usersCount", len(usersDomain)),
		zap.String("journey", "findAllUsers"))

	return usersDomain, nil
}
