package main

import (
	"three_dimensional/api"
	"three_dimensional/dao"
)

func main()  {
	dao.InitDB()
	api.InitEngine()
}
