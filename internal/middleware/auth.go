package middleware

import (
	"api_course/internal/user"

	"github.com/gin-gonic/gin"
)

func Authentication(userService *user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "unauthorize",
			})
			c.Abort()
		}
		c.Next()
	}
}