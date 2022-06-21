package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	route.GET("/exercise/:id", func(ctx *gin.Context) {})

	route.Run(":8080")
}