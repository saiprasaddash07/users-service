package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/controllers/v1/userServices"
	"github.com/saiprasaddash07/users-service/controllers/v1/utils"
	"github.com/saiprasaddash07/users-service/helpers/request"
	"github.com/saiprasaddash07/users-service/helpers/response"
	"github.com/saiprasaddash07/users-service/helpers/util"
)

func Register(c *gin.Context) {
	userFromContext, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	user := userFromContext.(*request.User)

	if err := utils.ValidateUserDetails(user); err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	_, err := userServices.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.CREATE_USER_MESSAGE,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))
}
