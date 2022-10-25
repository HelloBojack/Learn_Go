package main

import (
	"gk"
	"net/http"
)

func main() {
	r := gk.New()

	r.Get("/", func(c *gk.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gk</h1>")
	})

	r.Get("/hello", func(c *gk.Context) {
		// expect /hello?name=gkktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/hello/:name", func(c *gk.Context) {
		// expect /hello/gkktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *gk.Context) {
		c.JSON(http.StatusOK, gk.J{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
