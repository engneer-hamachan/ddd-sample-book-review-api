package middleware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var ClaimUserID string

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "SECRET_KEY Error"})
			c.Abort()
			return
		}

		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(secretKey)
			return b, nil
		})

		if err == nil {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims["user_id"])
			ClaimUserID = string(claims["user_id"].(string))
			c.Next()
		} else {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
