package converter

import (
	"github.com/marcofilho/go-crud-api/src/model"
	"github.com/marcofilho/go-crud-api/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
