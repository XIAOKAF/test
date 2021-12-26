package dao

import "database/sql"

var DB *sql.DB

func InitDB()  {
	db, err := sql.Open("mysql","root:123456@tcp(localhost:3306)/Three_dimensional?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	DB = db
}
