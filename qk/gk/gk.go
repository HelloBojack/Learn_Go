package gk

import (
	"net/http"
	"strings"
)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) Get(path string, handler HandlerFunc) {
	e.router.addRoute("GET", path, handler)
}

func (e *Engine) Post(path string, handler HandlerFunc) {
	e.router.addRoute("POST", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := NewContext(w, req)
	c.handlers = middlewares
	e.router.handle(c)
}
