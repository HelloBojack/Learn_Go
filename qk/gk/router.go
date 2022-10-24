package gk

import (
	"net/http"
)

type HandlerFunc func(*Context)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(path, 0)

	key := method + "-" + path
	r.handlers[key] = handler
}

// func (r *router) getRoute(method string, path string) {

// 	root,ok:=r.roots[method]
// 	if !ok {
// 		return
// 	}

// }

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
