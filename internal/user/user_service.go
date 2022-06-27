package user

import (
	"api_course/internal/domain"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	var registerUser domain.RegisterUser
	err := ctx.ShouldBind(&registerUser)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Input",
		})
		return
	}

	if registerUser.Name == "" {
		ctx.JSON(400, gin.H{
			"message": "field name required",
		})
		return
	}

	if registerUser.Email == "" {
		ctx.JSON(400, gin.H{
			"message": "field email required",
		})
		return
	}

	if registerUser.Password == "" {
		ctx.JSON(400, gin.H{
			"message": "field password required",
		})
		return
	}

	if len(registerUser.Password) < 6 {
		ctx.JSON(400, gin.H{
			"message": "field password must be more than 6 character",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	registerUser.Password = string(hashedPassword)

	user := domain.User{
		Name:     registerUser.Name,
		Email:    registerUser.Email,
		Password: registerUser.Password,
		NoHp:     registerUser.NoHp,
	}

	if err := us.database.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed when create user",
		})
		return
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed when create user",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"token": token,
	})
}

func (us UserService) LoginUser(ctx *gin.Context) {
	var currentUser domain.LoginUser

	err := ctx.ShouldBind(&currentUser)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Input",
		})
		return
	}

	var user domain.User
	err = us.database.Where("email = ?", currentUser.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid email / password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentUser.Password))
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid email/password",
		})
		return
	}
	token, err := generateJWT(user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed when get user",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

var signature = []byte("mySecretSignatureKey")

func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "goexp",
	}
	fmt.Println("claims: ", claims)
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signature)

	if err != nil {
		return "", err
	}

	return stringToken, err
}

func (us UserService) DecriptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return signature, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}
	if !parsedToken.Valid {
		return data, errors.New("token invalid")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}