package main

import (
	"be-blog/src/config"
	"be-blog/src/libs/logger"
	"be-blog/src/routers"
	"fmt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	app := iris.New()
	{
		db := config.InitDatabase()
		defer db.Close()
	}
	err := config.NewCloudinary()
	if err != nil {
		panic(err)
	}
	logger.InitLog(logger.Config{
		Folder: "logs/",
		Ext:    ".log",
	})
	app.HandleDir("/images", "./uploads")
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://vananh.com:3000", "http://vananh.com", "http://localhost:8000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)
	config.InitMail()
	routers.RegisterRouter(app)
	app.Listen(fmt.Sprintf(":%d", viper.GetInt("port")), iris.WithoutServerError(iris.ErrServerClosed))
}
