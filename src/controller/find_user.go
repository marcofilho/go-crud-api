package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/controller/model/response"
	"github.com/marcofilho/go-crud-api/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Initializing FindUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userID := c.Param("userID")

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		logger.Error("Error trying to validate userID", err,
			zap.String("userID", userID),
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError("Invalid userID format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(userID)
	if err != nil {
		logger.Error("Error trying to call findUserByID service", err,
			zap.String("userID", userID),
			zap.String("journey", "findUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("userID", userID),
		zap.String("journey", "findUserByID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Initializing findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	email := c.Param("email")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error trying to validate email", err,
			zap.String("email", email),
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError("User email is not in a valid format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call findByEmail service", err,
			zap.String("userID", email),
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("email", email),
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info("Initializing FindAllUsers controller",
		zap.String("journey", "findAllUsers"),
	)

	usersDomain, err := uc.service.FindAllUsers()
	if err != nil {
		logger.Error("Error trying to call findAllUsers service", err,
			zap.String("journey", "findAllUsers"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindAllUsers controller executed successfully",
		zap.String("journey", "findAllUsers"),
	)

	var listUsersResponse []response.UserResponse

	for _, userDomain := range usersDomain {
		listUsersResponse = append(listUsersResponse, view.ConvertDomainToResponse(userDomain))
	}

	c.JSON(http.StatusOK, listUsersResponse)
}
