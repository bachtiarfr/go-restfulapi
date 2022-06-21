package exercise

import (
	"api_course/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseService struct {
	db *gorm.DB
}

func NewExerciseService(database *gorm.DB) *ExerciseService {
	return &ExerciseService{
		db: database,
	}
}

func (ex ExerciseService) GetExercise(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = ex.db.Where("id = ?", id).Preload("Question").Take(&exercise).Error

	if err != nil {
		ctx.JSON(404, gin.H{
			"message" : "data not found",
		})
		return
	}
	ctx.JSON(200, exercise)

}