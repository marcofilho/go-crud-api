package converter

import (
	"github.com/marcofilho/go-crud-api/src/model"
	"github.com/marcofilho/go-crud-api/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID.String())

	return domain
}
