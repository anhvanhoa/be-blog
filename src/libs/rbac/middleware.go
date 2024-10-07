package rbac

import (
	"be-blog/src/constants"
	"be-blog/src/libs/errors"
	"be-blog/src/libs/jwt"
	"be-blog/src/libs/logger"
	"be-blog/src/services/session"
	"strings"

	"github.com/kataras/iris/v12"
)

func MiddlewarePermission(ctx iris.Context) {
	route := ctx.GetCurrentRoute()
	key := route.Method() + ":" + route.Path()
	router := routers[key]
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
	ok := checkPer(roles, route.Path(), route.Method())
	if !ok {
		err := errors.NewErrorUnauthorized("Bạn không có quyền truy cập")
		logger.Log(ctx, err)
		return
	}
	ctx.Next()
}

func checkPer(roles []string, path, method string) bool {
	for _, router := range routersPer {
		if router.Path == path && router.Method == method {
			for _, role := range roles {
				if role == router.NamePer {
					return true
				}
			}
			return false
		}
	}
	return false
}
