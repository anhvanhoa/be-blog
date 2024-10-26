package controllers

import (
	"be-blog/src/entities"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	tag_service "be-blog/src/services/tag"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func GetTags(ctx iris.Context) {
	tags, err := tag_service.GetTags()
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Data:    tags,
		Message: "Danh sách tag",
	})
}

func CreateTag(ctx iris.Context) {
	var tag models.TagReq
	err := ctx.ReadJSON(&tag)
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	if _, err := govalidator.ValidateStruct(tag); err != nil {
		logger.Log(ctx, err)
		return
	}

	err = tag_service.CreateTag(tag)
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Message: "Tạo tag thành công",
	})
}

func GetTag(ctx iris.Context) {
	id := ctx.Params().Get("id")
	if id == "create" {
		tag := entities.Tag{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		ctx.JSON(Response{
			Data:    tag,
			Message: "Lấy tag thành công",
		})
		return
	}
	tag, err := tag_service.GetTag(id)
	if err != nil {
		err := errors.NewError(err).NotFound()
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Data:    tag,
		Message: "Lấy tag thành công",
	})
}

func UpdateTag(ctx iris.Context) {
	id := ctx.Params().Get("id")
	var tag models.TagReq
	err := ctx.ReadJSON(&tag)
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	if _, err := govalidator.ValidateStruct(tag); err != nil {
		logger.Log(ctx, err)
		return
	}

	err = tag_service.UpdateTag(id, tag)
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Message: "Cập nhật tag thành công",
	})
}

func DeleteTag(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := tag_service.DeleteTag(id)
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Message: "Xóa tag thành công",
	})
}
