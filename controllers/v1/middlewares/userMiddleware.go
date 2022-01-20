package middlewares

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/controllers/v1/utils"
)

func GetRequestBodyUser(apiType string, registerRequiredFields []string, registerOptionalFields []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestObj interface{}

		if err := context.ShouldBind(&requestObj); err != nil {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		userJSON := requestObj.(map[string]interface{})

		user, ok := utils.ValidateAndParseUserFields(userJSON, registerRequiredFields, registerOptionalFields)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		if apiType == constants.API_TYPE_EDIT_USER {
			userId, err := strconv.ParseInt(context.Query("userId"), 10, 64)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"status":  constants.API_FAILED_STATUS,
					"message": constants.INVALID_USER_ID,
				})
				return
			}
			log.Println("userId:", userId)
			user.UserId = userId
		}

		if err := utils.ValidateUserDetails(user, apiType); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": err.Error(),
			})
			return
		}

		context.Set("user", user)
		context.Next()
	}
}
