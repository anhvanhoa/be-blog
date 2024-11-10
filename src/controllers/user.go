package controllers

import (
	"be-blog/src/libs/logger"
	"be-blog/src/services/user"

	"github.com/kataras/iris/v12"
)

func GetUsers(ctx iris.Context) {
	users, err := user.GetUsers()
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    users,
		Message: "Lấy danh sách user thành công",
	})
}

func GetUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	user, err := user.GetUser(id)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    user,
		Message: "Lấy user thành công",
	})
}
