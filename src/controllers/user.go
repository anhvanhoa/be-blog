package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	"be-blog/src/services/user"

	"github.com/asaskevich/govalidator"
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

func UpdateUserStatus(ctx iris.Context) {
	var req models.StatusUserReq
	if err := ctx.ReadJSON(&req); err != nil {
		err = errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		err = errors.NewError(err.(govalidator.Errors).Errors()[0]).BadRequest()
		logger.Log(ctx, err)
		return
	}
	if err := user.UpdateUserStatus(req); err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Cập nhật trạng thái user thành công",
	})
}

func UpdateUserRoles(ctx iris.Context) {
	var req models.RolesUserReq
	if err := ctx.ReadJSON(&req); err != nil {
		err = errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		err = errors.NewError(err.(govalidator.Errors).Errors()[0]).BadRequest()
		logger.Log(ctx, err)
		return
	}
	if err := user.UpdateUserRoles(req); err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Cập nhật roles user thành công",
	})
}
