package v1

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/controllers/v1/userServices"
	"github.com/saiprasaddash07/users-service/controllers/v1/utils"
	"github.com/saiprasaddash07/users-service/helpers/request"
	"github.com/saiprasaddash07/users-service/helpers/response"
	"github.com/saiprasaddash07/users-service/helpers/util"
)

func RegisterUserHandler(c *gin.Context) {
	userFromContext, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	user := userFromContext.(*request.User)

	if err := utils.ValidateUserDetails(user, constants.API_TYPE_CREATE_USER); err != nil {
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

func EditUserHandler(c *gin.Context) {
	userFromContext, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	user := userFromContext.(*request.User)

	userRes, err := userServices.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	createResponse := &response.UserEdit{
		UserId:    userRes.UserId,
		FirstName: userRes.FirstName,
		LastName:  userRes.LastName,
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.EDIT_USER_MESSAGE,
		Result:  createResponse,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))
}

func DeleteUserHandler(c *gin.Context) {
	userFromContext, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	user := userFromContext.(*request.User)

	err := userServices.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, util.SendSuccessResponse(constants.DELETE_USER_MESSAGE))
}

func FetchUserHandler(c *gin.Context) {
	userIds := c.Query("userIds")

	var userIdsArr []int

	// Parse user id from query string like this ?userIds=1,2,3,4,5
	if userIds != "" {
		userIdArr := strings.Split(userIds, ",")
		for _, userId := range userIdArr {
			id, err := strconv.Atoi(userId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": constants.USER_ID_NOT_ARRAY_OF_INT,
				})
				return
			}
			userIdsArr = append(userIdsArr, id)
		}
	}

	var args []interface{}

	if len(userIdsArr) > 0 {
		for _, user := range userIdsArr {
			args = append(args, user)
		}
		userRes, err := userServices.GetUser(c, args, "userId", 0, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, util.StructToJSON(userRes))
	} else {
		err := errors.New(constants.ERROR_NO_USER_EXIST)
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}
}