package gk

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup

	engine *Engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: g.engine,
	}
	g.engine.groups = append(g.engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middleware...)
}

func (g *RouterGroup) Get(path string, handler HandlerFunc) {
	g.engine.router.addRoute("GET", g.prefix+path, handler)
}

func (g *RouterGroup) Post(path string, handler HandlerFunc) {
	g.engine.router.addRoute("POST", g.prefix+path, handler)
}
