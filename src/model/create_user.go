package model

import (
	"fmt"

	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
