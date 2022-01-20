package DAO

import (
	"errors"
	"log"

	"github.com/saiprasaddash07/users-service/constants"
	"github.com/saiprasaddash07/users-service/helpers/request"
	"github.com/saiprasaddash07/users-service/services/db"
)

func CreateUser(user *request.User) error {
	rows, err := db.GetClient(constants.DB_WRITER).Exec("INSERT INTO users(firstName, lastName, password, email, mobileNo, gender) VALUES (?,?,?,?,?,?);", user.FirstName, user.LastName, user.Password, user.Email, user.MobileNo, user.Gender)
	if err != nil {
		return err
	}
	userId, err := rows.LastInsertId()
	if err != nil {
		return err
	}
	user.UserId = userId
	return nil
}

func DoesUserExist(user *request.User) error {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM users WHERE email=? OR mobileNo=?;", user.Email, user.MobileNo).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return nil
	}
	return errors.New(constants.ERROR_IN_STORING_UNIQUE_USER)
}

func UpdateUser(user *request.User) error {
	log.Println(user.UserId, user.LastName)
	_, err := db.GetClient(constants.DB_WRITER).Exec("UPDATE users SET firstName=?, lastName=? WHERE userId=?;", user.FirstName, user.LastName, user.UserId)
	if err != nil {
		return err
	}
	return nil
}

func GetPassword(user *request.User) (string, error) {
	var password string
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT password FROM users WHERE userId=?;", user.UserId).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func DeleteUser(user *request.User) error {
	_, err := db.GetClient(constants.DB_WRITER).Exec("UPDATE users SET isDeleted=? WHERE userId=?;", "true", user.UserId)
	if err != nil {
		return err
	}
	return nil
}