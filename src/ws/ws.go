package ws

import (
	"be-blog/src/constants"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func Handler() iris.Handler {
	hub := newHub()
	go hub.run()
	return func(c iris.Context) {

		id := c.GetCookie(constants.COOKIE_VISIT)
		if id == "" {
			id = uuid.New().String()
		}
		serveWs(hub, id, c.ResponseWriter(), c.Request())
	}
}

func RegisterWs(app *iris.Application) {
	app.Get("/ws", Handler())
}
