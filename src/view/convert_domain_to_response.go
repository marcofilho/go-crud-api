package view

import (
	"github.com/marcofilho/go-crud-api/src/controller/model/response"
	"github.com/marcofilho/go-crud-api/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}

}
