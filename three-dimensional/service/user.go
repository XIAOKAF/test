package service

import (
	"database/sql"
	"three_dimensional/dao"
	"three_dimensional/modle"
)

func IsRepeatUsername(username string) (bool, error) {

	_, err := dao.SelectUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Register(user modle.User) error {
	err := dao.CreateUser(user)
	return err
}

func Login(username, password string) (bool, error) {
	user, err := dao.Login(username)

	if err != nil {
		if password == user.Password {
			return true, err
		}
		return false, nil
	} else {
		if password == user.Password {
			return true, err
		}
		return false, nil
	}
}

