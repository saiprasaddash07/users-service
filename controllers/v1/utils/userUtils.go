package utils

import (
	"errors"
	"fmt"

	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/helpers/request"
	"github.com/saiprasaddash07/users-service/helpers/util"
)

func ValidateAndParseUserFields(userJSON map[string]interface{}, requiredFields []string, optionalFields []string) (*request.User, bool) {
	lengthDiffRequiredFieldsAndUserJSON := len(userJSON) - len(requiredFields)
	if lengthDiffRequiredFieldsAndUserJSON < 0 || len(optionalFields) < lengthDiffRequiredFieldsAndUserJSON {
		return nil, false
	}

	countOfReqFields := len(requiredFields)
	var user request.User
	for k, v := range userJSON {
		if util.Contains(requiredFields, k) {
			countOfReqFields--
		} else if !util.Contains(optionalFields, k) {
			return nil, false
		}

		valueType := fmt.Sprintf("%T", v)
		switch k {
		case "firstName":
			if valueType == "string" {
				user.FirstName = v.(string)
			} else {
				return &user, false
			}
		case "lastName":
			if valueType == "string" {
				user.LastName = v.(string)
			} else {
				return &user, false
			}
		case "password":
			if valueType == "string" {
				user.Password = v.(string)
			} else {
				return &user, false
			}
		case "email":
			if valueType == "string" {
				user.Email = v.(string)
			} else {
				return &user, false
			}
		case "mobileNo":
			if valueType == "string" {
				user.MobileNo = v.(string)
			} else {
				return &user, false
			}
		case "gender":
			if valueType == "float64" && util.IsInteger(v.(float64)) {
				user.Gender = int(v.(float64))
			} else {
				return &user, false
			}
		default:
			return nil, false
		}
	}
	if countOfReqFields == 0 {
		return &user, true
	}
	return nil, false
}

func ValidateUserDetails(user *request.User, apiType string) error {
	if apiType == constants.API_TYPE_CREATE_USER {
		if len(user.FirstName) < constants.MIN_LENGTH_OF_FIRSTNAME || len(user.FirstName) > constants.MAX_LENGTH_OF_FIRSTNAME {
			return errors.New(constants.INVALID_REQUEST)
		}

		if len(user.LastName) < constants.MIN_LENGTH_OF_LASTNAME || len(user.LastName) > constants.MAX_LENGTH_OF_LASTNAME {
			return errors.New(constants.INVALID_REQUEST)
		}

		if len(user.Email) == 0 {
			return errors.New(constants.INVALID_REQUEST)
		}
		email, ok := util.ValidEmail(user.Email)
		if !ok {
			return errors.New(constants.INVALID_MAIL_ID)
		}
		user.Email = email

		if len(user.Password) < constants.MIN_LENGTH_OF_PASSWORD || len(user.Password) > constants.MAX_LENGTH_OF_PASSWORD {
			return errors.New(constants.INVALID_REQUEST)
		}

		if len(user.MobileNo) != 10 {
			return errors.New(constants.INVALID_REQUEST)
		}
	} else if apiType == constants.API_TYPE_EDIT_USER {
		if len(user.FirstName) < constants.MIN_LENGTH_OF_FIRSTNAME || len(user.FirstName) > constants.MAX_LENGTH_OF_FIRSTNAME {
			return errors.New(constants.INVALID_REQUEST)
		}

		if len(user.LastName) < constants.MIN_LENGTH_OF_LASTNAME || len(user.LastName) > constants.MAX_LENGTH_OF_LASTNAME {
			return errors.New(constants.INVALID_REQUEST)
		}
	}

	return nil
}
