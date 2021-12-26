package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"three_dimensional/modle"
	"three_dimensional/service"
	"three_dimensional/tool"
)

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := modle.User{
		Username:         username,
		Password:         password,
	}

	flag, err := service.IsRepeatUsername(user.Username)

	if err != nil {
		fmt.Println(err)
	}

	if !flag && err == nil {
		err := service.Register(user)
		if err != nil {
			panic(err)
		}
		tool.Success(ctx)
	} else{
		tool.Failure(ctx)
	}
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := modle.User{
		Username:         username,
		Password:         password,
	}

	flag, err := service.Login(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	if flag != true {
		tool.Failure(ctx)

	}else{
		tool.Success(ctx)
	}
}
