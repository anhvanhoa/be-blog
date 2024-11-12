package controllers

import (
	"be-blog/src/libs/logger"
	"be-blog/src/services/statistic"

	"github.com/kataras/iris/v12"
)

func Statistic(ctx iris.Context) {
	res, err := statistic.Statistic()

	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Data:    res,
		Message: "Thống kê thành công",
	})
}

func StatisticCommentNew(ctx iris.Context) {
	res, err := statistic.StatisticCommentNew()

	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Data:    res,
		Message: "Thống kê bình luận mới thành công",
	})
}
