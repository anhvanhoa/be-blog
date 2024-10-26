package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	category_service "be-blog/src/services/category"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func GetCategories(ctx iris.Context) {
	categories, err := category_service.GetCategories(true)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    categories,
		Message: "Success",
	})
}

func GetManagerCategories(ctx iris.Context) {
	categories, err := category_service.GetAll()
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    categories,
		Message: "Lấy danh sách danh mục thành công",
	})
}

func GetCategoryById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	if id == "create" {
		category := models.Category{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		ctx.JSON(Response{
			Data:    category,
			Message: "Lấy danh mục thành công",
		})
		return
	}
	category, err := category_service.GetCategoryById(id)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    category,
		Message: "Lấy danh mục thành công",
	})
}

func CreateCategory(ctx iris.Context) {
	category := models.CategoryReq{}
	err := ctx.ReadJSON(&category)
	if err != nil {
		err := errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(category); err != nil {
		errs := err.(govalidator.Errors).Errors()
		logger.Log(ctx, errors.NewError(errs[0]).BadRequest())
		return
	}
	err = category_service.CreateCategory(category)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Tạo danh mục thành công",
	})
}

func UpdateCategory(ctx iris.Context) {
	id := ctx.Params().Get("id")
	category := models.CategoryReq{}
	err := ctx.ReadJSON(&category)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(category); err != nil {
		errs := err.(govalidator.Errors).Errors()
		logger.Log(ctx, errs[0])
		return
	}
	err = category_service.UpdateCategory(id, category)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Cập nhật danh mục thành công",
	})
}

func DeleteCategory(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := category_service.DeleteCategory(id)
	if err != nil {
		err := errors.NewError(err)
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Xóa danh mục thành công",
	})
}
