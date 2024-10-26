package rbac

import "github.com/kataras/iris/v12"

type FunPermission func(roles ...string) bool

type Route struct {
	app    iris.Party
	Path   string
	Method string
	Auth   bool
	Per    FunPermission
}

var routers = map[string]*Route{}
