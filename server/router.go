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
			userGroupV1.POST("/register", middlewares.GetRequestBodyUser(constants.API_TYPE_CREATE_USER, constants.USER_REGISTER_REQUIRED_FIELDS, constants.USER_REGISTER_OPTIONAL_FIELDS), v1.RegisterUserHandler)
			userGroupV1.POST("/edit", middlewares.GetRequestBodyUser(constants.API_TYPE_EDIT_USER, constants.USER_EDIT_REQUIRED_FIELDS, constants.USER_EDIT_OPTIONAL_FIELDS), v1.EditUserHandler)
			userGroupV1.POST("/delete", middlewares.GetRequestBodyUser(constants.API_TYPE_DELETE_USER, constants.USER_DELETE_REQUIRED_FIELDS, constants.USER_DELETE_OPTIONAL_FIELDS), v1.DeleteUserHandler)
			userGroupV1.POST("/fetch", v1.FetchUserHandler)
		}
	}

	return router
}
