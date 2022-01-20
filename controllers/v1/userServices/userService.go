package userServices

import (
	"errors"
	"strings"

	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/helpers/DAO"
	"github.com/saiprasaddash07/users-service/helpers/request"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *request.User) (*request.User, error) {
	user.Password = strings.TrimSpace(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(constants.ERROR_IN_HASHING_PASSWORD)
	}
	user.Password = string(hashedPassword)

	if err := DAO.DoesUserExist(user); err != nil {
		return nil, errors.New(constants.ERROR_IN_STORING_UNIQUE_USER)
	}

	if err := DAO.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
