package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Render index.html
	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
