package rbac

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

type Config struct {
	RelativePath string
}

func NewRoute(app iris.Party, config *Config) *Route {
	router := app.Party(config.RelativePath)
	return &Route{
		app: router,
	}
}

func AddRoute(r *Route) {
	key := r.Method + ":" + r.Path
	routers[key] = r
}

func (r Route) Get(Path string, Per FunPermission, Auth bool, Handlers ...iris.Handler) {
	r.Method = http.MethodGet
	r.Path = r.app.GetRelPath() + Path
	r.Auth = Auth
	r.Per = Per
	r.app.Get(Path, Handlers...)
	AddRoute(&r)
}

func (r Route) Post(Path string, Per FunPermission, Auth bool, Handlers ...iris.Handler) {
	r.Method = http.MethodPost
	r.Path = r.app.GetRelPath() + Path
	r.Per = Per
	r.Auth = Auth
	r.app.Post(Path, Handlers...)
	AddRoute(&r)
}

func (r Route) Put(Path string, Per FunPermission, Auth bool, Handlers ...iris.Handler) {
	r.Method = http.MethodPut
	r.Path = r.app.GetRelPath() + Path
	r.Per = Per
	r.Auth = Auth
	r.app.Put(Path, Handlers...)
	AddRoute(&r)
}

func (r Route) Delete(Path string, Per FunPermission, Auth bool, Handlers ...iris.Handler) {
	r.Method = http.MethodDelete
	r.Path = r.app.GetRelPath() + Path
	r.Per = Per
	r.Auth = Auth
	r.app.Delete(Path, Handlers...)
	AddRoute(&r)
}

func (r Route) Patch(Path string, Per FunPermission, Auth bool, Handlers ...iris.Handler) {
	r.Method = http.MethodPatch
	r.Path = r.app.GetRelPath() + Path
	r.Per = Per
	r.Auth = Auth
	r.app.Patch(Path, Handlers...)
	AddRoute(&r)
}
