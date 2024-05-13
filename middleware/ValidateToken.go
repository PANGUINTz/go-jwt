package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthVerify() gin.HandlerFunc {
	return func (c* gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_TOKEN"))
		header := c.Request.Header.Get("Authorization")
		result := strings.Replace(header, "Bearer ", "", -1)
		token, err := jwt.Parse(result, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil,  fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if err != nil || token == nil {
			c.JSON(http.StatusForbidden, gin.H{"status": false, "message": "Invalid token"})
			c.Abort()
			return
		}
		
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
			fmt.Print(claims)
			} else {
				c.JSON(http.StatusForbidden, gin.H{"status":false, "message": err.Error()})
				fmt.Print(claims)
				c.Abort()
		}
		c.Next()
	}
}