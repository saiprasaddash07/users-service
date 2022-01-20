package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/users-service/constants"
	v1 "github.com/saiprasaddash07/users-service/controllers/v1"
	"github.com/saiprasaddash07/users-service/controllers/v1/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	version1 := router.Group("api/v1")
	{
		userGroupV1 := version1.Group("user")
		{
			userGroupV1.POST("/register", middlewares.GetRequestBodyUser(constants.HTTP_METHOD_POST, constants.USER_REGISTER_REQUIRED_FIELDS, constants.USER_REGISTER_OPTIONAL_FIELDS), v1.Register)
		}
	}

	return router
}
