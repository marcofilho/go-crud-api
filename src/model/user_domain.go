package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) EncryptPassword() {
	has := md5.New()
	defer has.Reset()
	has.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(has.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	FindUserByID(id string) (*userDomain, *rest_err.RestErr)
	UpdateUser(id string) *rest_err.RestErr
	FindUserByEmail(email string) (*userDomain, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
