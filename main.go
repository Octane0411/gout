package main

import (
	"github.com/octane0411/gout"
	"net/http"
)

func main() {
	r := gout.New()
	r.GET("/index", func(c *gout.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gout.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gout.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gout.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gout.Context) {
			c.JSON(http.StatusOK, gout.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	r.Run(":8000")
}
