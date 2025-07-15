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
	logger.Info("Initializing UpdateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest
	userID := c.Param("userID")

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		logger.Error("Error trying to validate userID", err,
			zap.String("userID", userID),
			zap.String("journey", "updateUser"),
		)
		errorMessage := rest_err.NewBadRequestError("Invalid userID format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userID, domain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userID", userID),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
