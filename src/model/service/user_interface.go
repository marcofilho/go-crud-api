package service

import (
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(id string, model model.UserDomainInterface) *rest_err.RestErr
	FindUserByID(id string) (*model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
