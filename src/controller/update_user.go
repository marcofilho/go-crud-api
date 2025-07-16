package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-crud-api/src/configuration/logger"
	"github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/configuration/validation"
	"github.com/marcofilho/go-crud-api/src/controller/model/request"
	"github.com/marcofilho/go-crud-api/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userID := c.Param("userID")
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		logger.Error("Error trying to validate userID", err,
			zap.String("userID", userID),
			zap.String("journey", "updateUser"),
		)
		errRest := rest_err.NewBadRequestError("Invalid userID, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userID, domain)
	if err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"updateUser controller executed successfully",
		zap.String("userID", userID),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
