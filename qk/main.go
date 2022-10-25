package main

import (
	"gk"
	"log"
	"net/http"
	"time"
)

func Logger() gk.HandlerFunc {
	return func(c *gk.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
func onlyForV2() gk.HandlerFunc {
	return func(c *gk.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gk.New()

	r.Use(Logger())
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

	v1 := r.Group("/v1")

	{
		v1.Get("/", func(c *gk.Context) {
			c.HTML(http.StatusOK, "<h1>Hello v1</h1>")
		})

		v1.Get("/hello", func(c *gk.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.Get("/hello/:name", func(c *gk.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
