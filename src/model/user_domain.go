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
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	has := md5.New()
	defer has.Reset()
	has.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(has.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(id string) *rest_err.RestErr
	FindUserByID(id string) (*UserDomain, *rest_err.RestErr)
	FindUserByEmail(email string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(id string) *rest_err.RestErr
}
