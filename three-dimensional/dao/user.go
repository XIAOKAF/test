package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"three_dimensional/modle"
)

func SelectUsername(username string) (modle.User, error) {
	user := modle.User{}

	tx, _ := DB.Begin()

	row := DB.QueryRow("select id,username from Users where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
		fmt.Println(row.Err())
	}

	err := row.Scan(&user.Id, &user.Username)
	if err != nil {
		return user, err
	}
	tx.Commit()
	return user, nil
}

func CreateUser(user modle.User) error {
	_, err := DB.Exec("INSERT INTO Users(username,password)"+"values (?,?,?,?);", user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Login(username string) (modle.User, error) {
	user := modle.User{}

	row := DB.QueryRow("select password from Users where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
		fmt.Println(row.Err())
	}

	err := row.Scan(&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
