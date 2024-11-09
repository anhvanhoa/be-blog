package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	"be-blog/src/services/account"
	media_service "be-blog/src/services/media"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func Me(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)
	profile, err := account.Profile(data.ID)
	if err != nil {
		logger.Log(ctx, err)
	}
	data.Avatar = profile.Avatar
	ctx.JSON(Response{
		Data:    data,
		Message: "Hello",
	})
}

func Profile(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)
	profile, err := account.Profile(data.ID)
	if err != nil {
		logger.Log(ctx, err)
	}
	ctx.JSON(Response{
		Data:    profile,
		Message: "Lấy thông tin thành công",
	})
}

func UpdateProfile(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)
	var profile models.ProfileUpdateReq
	err := ctx.ReadJSON(&profile)
	if err != nil {
		err := errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
	}

	if _, err := govalidator.ValidateStruct(profile); err != nil {
		logger.Log(ctx, err.(govalidator.Errors).Errors()[0])
		return
	}

	profile.ID = data.ID
	err = account.UpdateProfile(profile, data)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Cập nhật thông tin thành công",
	})
}

func UploadAvatar(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)
	file, info, err := ctx.FormFile("file")
	if err != nil {
		logger.Log(ctx, err)
	}
	url, err := media_service.UploadImage(file, info, data.ID, info.Filename)
	if err != nil {
		logger.Log(ctx, err)
	}
	err = account.UploadAvatar(data.ID, url)
	if err != nil {
		logger.Log(ctx, err)
	}
	ctx.JSON(Response{
		Data:    url,
		Message: "Upload ảnh đại diện thành công",
	})
}

func DeleteAvatar(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)
	err := account.DeleteAvatar(data.ID)
	if err != nil {
		logger.Log(ctx, err)
	}
	ctx.JSON(Response{
		Message: "Xóa ảnh đại diện thành công",
	})
}

func ChangePassword(ctx iris.Context) {
	user := ctx.Values().Get("user")
	data := user.(*jwt.PayloadUser)
	var req models.ChangePasswordReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		err := errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		logger.Log(ctx, err.(govalidator.Errors).Errors()[0])
		return
	}

	err = account.ChangePassword(data.ID, req)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Đổi mật khẩu thành công",
	})
}
