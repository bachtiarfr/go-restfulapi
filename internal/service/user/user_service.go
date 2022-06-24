package user

import (
	"api_course/internal/domain"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	database *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		database: db,
	}
}

func (us UserService) RegisterUser(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "invalid input",
		})
		return
	}

	if user.Name == "" {
		ctx.JSON(400, gin.H{
			"message" : "invalid input",
		})
		return
	}

	if user.Email == "" {
		ctx.JSON(400, gin.H{
			"message" : "invalid input",
		})
		return
	}

	if user.Password == "" {
		ctx.JSON(400, gin.H{
			"message" : "invalid input",
		})
		return
	}

	if len(user.Password) < 6 {
		ctx.JSON(400, gin.H{
			"message" : "password must more than 6 charracter",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err = us.database.Create(&user).Error

	if err := us.database.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"message" : "failed when create new user",
		})
		return
	}
	
}