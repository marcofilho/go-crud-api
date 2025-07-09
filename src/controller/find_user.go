package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Initializing FindUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userID := c.Param("userID")

	if _, err := uuid.Parse(userID); err != nil {
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
		logger.Error("Error in FindUserByID controller", err,
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
		logger.Error("Error trying to validate userID", err,
			zap.String("email", email),
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError("User email is not in a valid format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error in FindUserByEmail controller", err,
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
