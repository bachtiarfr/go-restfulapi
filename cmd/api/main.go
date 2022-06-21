package main

import (
	"api_course/internal/database"
	"api_course/internal/service/exercise"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// connect to database
	db := database.NewDatabaseConnection()
	exerciseService := exercise.NewExerciseService(db)

	route.GET("/exercise/:id", exerciseService.GetExercise)

	route.Run(":8080")
}