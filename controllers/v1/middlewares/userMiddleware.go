package middlewares

import (
	"net/http"

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

		if err := utils.ValidateUserDetails(user); err != nil {
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
