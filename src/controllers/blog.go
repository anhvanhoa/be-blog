package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	blog_service "be-blog/src/services/blog"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func GetBlogs(ctx iris.Context) {
	tag := ctx.URLParam("tag")
	blogs, err := blog_service.GetBlogs(tag)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    blogs,
		Message: "Success",
	})
}

func GetBlogBySlug(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	blog, err := blog_service.GetBlogBySlug(slug)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    blog,
		Message: "Success",
	})
}

func GetBlogById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	if id == "create" {
		create := new(models.BlogByID)
		create.Tags = []models.TagBlog{}
		ctx.JSON(Response{
			Data:    create,
			Message: "Success",
		})
		return
	}

	blog, err := blog_service.GetBlogById(id)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    blog,
		Message: "Success",
	})
}

func GetBlogsByCategory(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	blogs, err := blog_service.GetBlogsByCategory(slug)
	if err != nil {
		logger.Log(ctx, err)
		return
	}

	ctx.JSON(Response{
		Data:    blogs,
		Message: "Success",
	})
}

func GetManagerBlogs(ctx iris.Context) {
	blogs, err := blog_service.GetManagerBlogs()
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    blogs,
		Message: "Success",
	})
}

func CreateBlog(ctx iris.Context) {
	var blog models.BlogReq
	if err := ctx.ReadJSON(&blog); err != nil {
		err := errors.NewErrorBadRequest("Không thể đọc dữ liệu")
		logger.Log(ctx, err)
		return
	}
	var user = ctx.Values().Get("user")
	blog.AuthorId = user.(*jwt.PayloadUser).ID
	if _, err := govalidator.ValidateStruct(blog); err != nil {
		err := errors.NewError(err.(govalidator.Errors).Errors()[0]).BadRequest()
		logger.Log(ctx, err)
		return
	}
	blog.Status = new(bool)
	err := blog_service.CreateBlog(blog)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Success",
	})
}

func UpdateBlog(ctx iris.Context) {
	var blog models.BlogReq
	if err := ctx.ReadJSON(&blog); err != nil {
		err := errors.NewErrorBadRequest("Không thể đọc dữ liệu")
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(blog); err != nil {
		err := errors.NewError(err.(govalidator.Errors).Errors()[0]).BadRequest()
		logger.Log(ctx, err)
		return
	}
	user := ctx.Values().Get("user")
	blog.AuthorId = user.(*jwt.PayloadUser).ID
	err := blog_service.UpdateBlog(blog.ID, blog)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Success",
	})
}

func DeleteBlog(ctx iris.Context) {
	id := ctx.Params().Get("id")
	err := blog_service.DeleteBlog(id)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Xóa bài viết thành công",
	})
}
