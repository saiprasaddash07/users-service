package userServices

import (
	"context"
	"errors"
	"strings"

	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/controllers/v1/utils"
	"github.com/saiprasaddash07/users-service/helpers/DAO"
	"github.com/saiprasaddash07/users-service/helpers/request"
	"github.com/saiprasaddash07/users-service/helpers/response"
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

func UpdateUser(user *request.User) (*request.User, error) {
	if err := DAO.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(user *request.User) error {
	password, err := DAO.GetPassword(user)
	if err != nil {
		return err
	}

	if err := utils.Authenticate(password, user.Password); err != nil {
		return errors.New(constants.ERROR_IN_AUTHENTICATING_USER)
	}

	if err := DAO.DeleteUser(user); err != nil {
		return err
	}

	return nil
}

func GetUser(ctx context.Context, args []interface{}, column string, size int, nextToken int) ([]response.User, error) {
	ddBlog, err := DAO.GetUser(ctx, column, args, size, nextToken)
	if err != nil {
		return nil, err
	}
	return ddBlog, nil
}
