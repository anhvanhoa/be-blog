package controllers

import (
	"be-blog/src/libs/logger"
	"be-blog/src/services/mail"

	"github.com/kataras/iris/v12"
)

func HistoriesMail(ctx iris.Context) {
	result, err := mail.HistoriesMail()
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Data:    result,
		Message: "Lịch sử gửi mail",
	})
}
