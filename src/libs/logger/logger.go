package logger

import (
	"be-blog/src/libs/errors"
	"net/http"

	"github.com/kataras/iris/v12"
)

func Log(ctx iris.Context, err error) {
	logger.SetError(err)
	switch e := err.(type) {
	case *errors.Error:
		ctx.StatusCode(e.Status)
		ctx.JSON(iris.Map{"message": e.Message})
	default:
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Internal server error"})
	}
}
