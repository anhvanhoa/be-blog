package controllers

import (
	"be-blog/src/libs/jwt"

	"github.com/kataras/iris/v12"
)

func Me(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)

	ctx.JSON(Response{
		Data:    data,
		Message: "Hello",
	})
}
