package controllers

import (
	"be-blog/src/libs/logger"
	blog_service "be-blog/src/services/blog"

	"github.com/kataras/iris/v12"
)

func GetBlogs(ctx iris.Context) {
	blogs, err := blog_service.GetBlogs()
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
