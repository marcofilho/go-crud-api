package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/validation"
	"github.com/marcofilho/go-crud-api/src/controller/model/request"
	"github.com/marcofilho/go-crud-api/src/model"
	"github.com/marcofilho/go-crud-api/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Initializing CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userID", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
