package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	resterr "github.com/marcofilho/go-crud-api/src/configuration/rest_err"
	"github.com/marcofilho/go-crud-api/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("Invalid user request, error: %s", err))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
