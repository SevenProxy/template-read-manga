package middleware

import (
	"net/http"
	"readmanga-api-auth/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

const UserEmailKey = "user_email"

func AuthMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  0,
				"message": "Unauthorized: Authorization header ausente",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "7proxy" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  0,
				"message": "Unauthorized: Formato do token inválido",
			})
			return
		}

		token := parts[1]
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
