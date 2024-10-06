package main

import (
	"be-blog/src/config"
	"be-blog/src/libs/logger"
	"be-blog/src/routers"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

func main() {
	app := iris.New()
	{
		db := config.InitDatabase()
		defer db.Close()
	}
	logger.InitLog(logger.Config{
		Folder: "logs/",
		Ext:    ".log",
	})
	config.InitMail()
	routers.RegisterRouter(app)
	app.Listen(fmt.Sprintf(":%d", viper.GetInt("port")))
}
