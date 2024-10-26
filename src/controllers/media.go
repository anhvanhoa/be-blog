package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	media_service "be-blog/src/services/media"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func GetImages(ctx iris.Context) {
	user := ctx.Values().Get("user")
	userId := user.(*jwt.PayloadUser).ID
	photos, err := media_service.GetImages(userId)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    photos,
		Message: "Get images successfully",
	})
}

func UploadImage(ctx iris.Context) {
	user := ctx.Values().Get("user")
	userId := user.(*jwt.PayloadUser).ID
	title := ctx.FormValue("title")
	file, info, err := ctx.FormFile("file")
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	defer file.Close()
	err = media_service.UploadImage(file, info, userId, title)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Upload image successfully",
	})
}

func DeleteImage(ctx iris.Context) {
	user := ctx.Values().Get("user")
	userId := user.(*jwt.PayloadUser).ID
	id := ctx.Params().Get("id")
	err := media_service.DeleteImage(userId, id)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Delete image successfully",
	})
}

func UpdateImage(ctx iris.Context) {
	var photo models.PhotoReq
	if err := ctx.ReadJSON(&photo); err != nil {
		err := errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(photo); err != nil {
		errs := err.(govalidator.Errors).Errors()
		logger.Log(ctx, errs[0])
		return
	}
	user := ctx.Values().Get("user")
	userId := user.(*jwt.PayloadUser).ID
	photo.AuthorId = userId
	err := media_service.UpdateImage(photo)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Update image successfully",
	})
}
