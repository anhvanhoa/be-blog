package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerAuthRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/auth"})
	router.Post("/register", "Đăng ký tài khoản", false, controllers.Register)
	router.Get("/verify-account", "Kiểm tra tài khoản xác thực", false, controllers.CheckVerifyEmail)
	router.Post("/verify-account", "Xác thực tài khoản", false, controllers.VerifyEmail)
	router.Get("/resend-verify-email", "Gửi lại mã xác thực", false, controllers.ResendCodeVerifyEmail)
	router.Post("/login", "Đăng nhập", false, controllers.Login)
	// router.Post("/logout", logout)
	// router.Post("/refresh", refresh)
}
