package main

import (
	"github.com/octane0411/gout"
	"net/http"
)

func main() {
	r := gout.New()
	r.GET("/", func(c *gout.Context) {
		c.HTML(200, "<h1>Hello Gout</h1>")
	})

	r.GET("/hello", func(c *gout.Context) {
		c.String(200, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gout.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gout.Context) {
		c.JSON(http.StatusOK, gout.H{"filepath": c.Param("filepath")})
	})
	r.Run(":8000")
}
