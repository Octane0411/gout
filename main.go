package main

import (
	"github.com/octane0411/gout"
)

func main() {
	r := gout.New()
	r.GET("/", func(c *gout.Context) {
		c.HTML(200, "<h1>Hello Gout</h1>")
	})

	r.GET("/hello", func(c *gout.Context) {
		c.String(200, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gout.Context) {
		c.JSON(200, gout.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":8000")
}
