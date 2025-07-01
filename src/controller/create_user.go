package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcofilho/go-crud-api/src/configuration/validation"
	"github.com/marcofilho/go-crud-api/src/controller/model/request"
	"github.com/marcofilho/go-crud-api/src/controller/model/response"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal user request, error: %s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userResponse := response.UserResponse{
		ID:    uuid.NewString(),
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, &userResponse)
}
