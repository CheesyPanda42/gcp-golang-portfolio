package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.Static("/assets", "./assets")
	router.Static("/css", "./css")
	// Render index.html
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
