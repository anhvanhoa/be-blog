package routers

import (
	"be-blog/src/controllers"
	"be-blog/src/libs/rbac"

	"github.com/kataras/iris/v12/core/router"
)

func registerAuthRouter(app router.Party) {
	router := rbac.NewRoute(app, &rbac.Config{RelativePath: "/auth"})
	router.Post("/register", rbac.AllowAll(), false, controllers.Register)                        //Đăng ký tài khoản
	router.Get("/verify-account", rbac.AllowAll(), false, controllers.CheckVerifyEmail)           //Kiểm tra tài khoản xác thực
	router.Post("/verify-account", rbac.AllowAll(), false, controllers.VerifyEmail)               //Xác thực tài khoản
	router.Get("/resend-verify-email", rbac.AllowAll(), false, controllers.ResendCodeVerifyEmail) //Gửi lại mã xác thực
	router.Post("/login", rbac.AllowAll(), false, controllers.Login)                              //Đăng nhập
	router.Post("/logout", rbac.AllowAll(), true, controllers.Logout)                             //Đăng xuất
	router.Post("/forgot-password", rbac.AllowAll(), false, controllers.ForgotPassword)           //Quên mật khẩu
	router.Get("/reset-password", rbac.AllowAll(), false, controllers.CheckResetPassword)         //Đặt lại mật khẩu
	router.Post("/reset-password", rbac.AllowAll(), false, controllers.ResetPassword)             //Đặt lại mật khẩu
}
