package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(id string)

	GetJSONValeu() (string, error)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJSONValeu() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println("Error marshalling user domain:", err)
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) EncryptPassword() {
	has := md5.New()
	defer has.Reset()
	has.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(has.Sum(nil))
}
