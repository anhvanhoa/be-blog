package rbac

import "github.com/kataras/iris/v12"

type Route struct {
	app    iris.Party
	Path   string
	Method string
	Name   string
	Auth   bool
}

type RoutePer struct {
	Path    string
	Method  string
	Name    string
	Auth    bool
	NamePer string
}

var routers = map[string]*Route{}

var routersPer = []RoutePer{}
