package main

import (
	"github.com/octane0411/gout"
	"net/http"
)

func main() {
	r := gout.Default()
	r.GET("/", func(c *gout.Context) {
		c.String(http.StatusOK, "Hello\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gout.Context) {
		names := []string{"a"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
