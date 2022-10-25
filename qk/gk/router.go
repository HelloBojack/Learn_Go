package gk

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func parsePath(path string) []string {
	temp_parts := strings.Split(path, "/")
	parts := make([]string, 0)
	for _, item := range temp_parts {
		if item != "" {
			parts = append(parts, item)
		}
	}
	return parts
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

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePath(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	node := root.search(path, 0)

	if node != nil {
		parts := parsePath(node.path)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)

	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.path
		// if handler, ok := r.handlers[key]; ok {
		// 	handler(c)
		// }
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
