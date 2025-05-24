package middleware

import (
	"fmt"
	"net/http"
	"readmanga-api-auth/internal/auth"

	"github.com/gin-gonic/gin"
)

const UserEmailKey = "user_email"

func AuthMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		fmt.Println(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  0,
				"message": "Unauthorized: Token ausente",
			})
			return
		}
		fmt.Println(token)
		claims, err := auth.ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  0,
				"message": "Unauthorized: Token inválido",
			})
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  0,
				"message": "Unauthorized: Email inválido no token",
			})
			return
		}

		c.Set(UserEmailKey, email)
		c.Next()
	}
}
