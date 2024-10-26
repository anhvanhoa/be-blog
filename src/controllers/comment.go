package controllers

import (
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	"be-blog/src/services/comment"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
)

func GetCommentsByBlog(ctx iris.Context) {
	slug := ctx.Params().Get("slug")                 //Lấy slug từ URL
	comments, err := comment.GetComments(slug, true) //Lấy danh sách bình luận hoạt động
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    comments,
		Message: "Success",
	})
}

func CommentPost(ctx iris.Context) {
	var cmt models.CommentReq
	var user = ctx.Values().Get("user")
	if user.(*jwt.PayloadUser).ID != "" {
		cmt.AuthorId = user.(*jwt.PayloadUser).ID
	}
	if err := ctx.ReadJSON(&cmt); err != nil {
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(cmt); err != nil {
		err = errors.NewErrorBadRequest("Dữ liệu không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	authorComment := models.AuthorComment{
		ID:       cmt.AuthorId,
		Avatar:   user.(*jwt.PayloadUser).Avatar,
		FullName: user.(*jwt.PayloadUser).FullName,
		Username: user.(*jwt.PayloadUser).UserName,
	}
	err := comment.CreateComment(cmt, cmt.UsernameReceiver, authorComment)
	if err != nil {
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Bình luận thành công",
	})
}
