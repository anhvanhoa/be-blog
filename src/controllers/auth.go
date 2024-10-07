package controllers

import (
	"be-blog/src/constants"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/models"
	"be-blog/src/services/auth"
	"be-blog/src/services/session"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func Register(ctx iris.Context) {
	var req models.RegisterReq
	if err := ctx.ReadJSON(&req); err != nil {
		err = errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		err = errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	if req.Password != req.ConfirmPassword {
		err := errors.NewErrorBadRequest("Mật khẩu không khớp")
		logger.Log(ctx, err)
		return
	}

	result, err := auth.Register(req)
	if err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Data:    result,
		Message: "Đăng ký thành công",
	})
}

func CheckVerifyEmail(ctx iris.Context) {
	tokenEmail := ctx.URLParam("t")
	if tokenEmail == "" {
		err := errors.NewErrorBadRequest("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	_, ok := auth.CheckVerifyEmail(tokenEmail)
	if !ok {
		err := errors.NewErrorBadRequest("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Tài khoản tiến hành xác thực",
	})
}

func VerifyEmail(ctx iris.Context) {
	tokenEmail := ctx.URLParam("t")
	if tokenEmail == "" {
		err := errors.NewErrorBadRequest("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	email, ok := auth.CheckVerifyEmail(tokenEmail)
	if !ok {
		err := errors.NewErrorBadRequest("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	var req models.VerifyEmailReq
	if err := ctx.ReadJSON(&req); err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	err := auth.VerifyAccount(email, req.Code)
	if err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Xác thực tài khoản thành công",
	})
}

func ResendCodeVerifyEmail(ctx iris.Context) {
	tokenEmail := ctx.URLParam("t")
	if tokenEmail == "" {
		err := errors.NewErrorBadRequest("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	email, ok := auth.CheckVerifyEmail(tokenEmail)
	if !ok {
		err := errors.NewErrorBadRequest("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	err := auth.ResendCodeVerifyEmail(email)
	if err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	ctx.JSON(Response{
		Message: "Gửi mã xác thực thành công",
	})
}

func Login(ctx iris.Context) {
	var req models.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	if _, err := govalidator.ValidateStruct(req); err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	result, token, err := auth.Login(req)
	if err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	ip := ctx.RemoteAddr()
	err = session.CreateSession(models.Session{
		UserID:    result.ID,
		Token:     token,
		IP:        ip,
		ExpiredAt: jwt.CreateExpTime(constants.SIX_MONTH).Unix(),
	})
	if err != nil {
		err := errors.NewError(err).BadRequest()
		logger.Log(ctx, err)
		return
	}
	ctx.SetCookieKV(constants.COOKIE_AUTH, token, iris.CookieHTTPOnly(true), iris.CookiePath("/"), CookieSecure())
	ctx.JSON(Response{
		Data:    result,
		Message: "Đăng nhập thành công",
	})
}

func CookieSecure() iris.CookieOption {
	return func(_ *context.Context, c *http.Cookie, op uint8) {
		c.Secure = true
	}
}
