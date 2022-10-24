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
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Post("/login", func(c *gk.Context) {
		c.JSON(http.StatusOK, gk.J{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
