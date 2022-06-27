package main

import (
	"api_course/internal/database"
	"api_course/internal/exercise"
	"api_course/internal/middleware"
	"api_course/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	
	// init routes
	route := gin.Default()

	// connect to database
	db := database.NewDatabaseConnection()

	exerciseService := exercise.NewExerciseService(db)
	userService := user.NewUserService(db)

	route.GET("/exercise/:id", exerciseService.GetExercise)
	route.POST("/register", userService.RegisterUser)
	route.POST("/login", userService.LoginUser)
	
	route.POST("/exercises", middleware.Authentication(userService), exerciseService.CreateExercise)
	route.GET("/exercises/:exerciseId", middleware.Authentication(userService), exerciseService.GetExercise)
	route.GET("/exercises/:exerciseId/score", middleware.Authentication(userService), exerciseService.GetUserScore)
	route.POST("/exercises/:exerciseId/questions", middleware.Authentication(userService), exerciseService.CreateQuestions)
	route.POST("/exercises/:exerciseId/questions/:questionId/answer", middleware.Authentication(userService), exerciseService.CreateAnswer)

	route.Run(":8080")
}