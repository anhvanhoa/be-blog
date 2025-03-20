package rbac

import (
	"be-blog/src/constants"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/services/session"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func MiddlewarePermission(ctx iris.Context) {
	id := ctx.GetCookie(constants.COOKIE_VISIT)
	if id == "" {
		id, _ := uuid.NewV7()
		ctx.SetCookieKV(constants.COOKIE_VISIT, id.String(), iris.CookieHTTPOnly(true), iris.CookiePath("/"), iris.CookieSameSite(http.SameSiteNoneMode), CookieSecure(), iris.CookieDomain(ctx.Domain()))
	}
	route := ctx.GetCurrentRoute()
	key := route.Method() + ":" + route.Path()
	router := routers[key]
	if router == nil {
		err := errors.NewErrorBadRequest("Không tìm thấy route")
		logger.Log(ctx, err)
		return
	}
	if !router.Auth {
		ctx.Next()
		return
	}
	idToken := ctx.GetCookie(constants.COOKIE_AUTH)
	if idToken == "" {
		err := errors.NewErrorUnauthorized("Bạn chưa đăng nhập")
		logger.Log(ctx, err)
		return
	}
	sessionLogin, err := session.GetSessionByToken(idToken)
	if err != nil {
		err := errors.NewError(err).SetStatus(iris.StatusUnauthorized)
		logger.Log(ctx, err)
		return
	}
	if sessionLogin.Token != idToken || sessionLogin.Token == "" {
		err := errors.NewErrorUnauthorized("Token không hợp lệ")
		logger.Log(ctx, err)
		return
	}
	// Check token
	payloadToken, err := jwt.ParseTokenPayloadUser(idToken)
	if err != nil {
		err := errors.NewError(err).SetStatus(iris.StatusUnauthorized)
		logger.Log(ctx, err)
		return
	}
	roles := strings.Split(payloadToken.Roles, ",")
	isAdmin := checkAdmin(roles...)
	if isAdmin {
		ctx.Values().Set("user", payloadToken)
		ctx.Next()
		return
	}
	ok := router.Per(roles...)
	if !ok {
		err := errors.NewErrorUnauthorized("Bạn không có quyền truy cập")
		logger.Log(ctx, err)
		return
	}
	ctx.Values().Set("user", payloadToken)
	ctx.Next()
}

func checkAdmin(roles ...string) bool {
	for _, role := range roles {
		if strings.EqualFold(role, RoleAdmin) {
			return true
		}
	}
	return false
}

func CookieSecure() iris.CookieOption {
	return func(_ *context.Context, c *http.Cookie, op uint8) {
		c.Secure = true
	}
}
