package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return func(c *gin.Context) {
		raw := c.GetHeader("Authorization")
		if !strings.HasPrefix(raw, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(strings.TrimPrefix(raw, "Bearer "), func(t *jwt.Token) (interface{}, error) { return secret, nil })
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("uid", uint(claims["uid"].(float64)))
		c.Next()
	}
}
