package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (app *application) AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context)   {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized,  gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}

		token , err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _ , ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(app.jwtSecret), nil
		})

		if (!token.Valid || err != nil) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims , ok  := token.Claims.(jwt.MapClaims)
		if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

		userId := claims["userId"].(float64)

		user, err :=  app.models.Users.Get(int(userId))
		if err != nil {
			 c.JSON(http.StatusNotFound, gin.H{"error": "Unauthorized access"})
            c.Abort()
            return
		}

		c.Set("user", user)
		c.Next()
	}
}