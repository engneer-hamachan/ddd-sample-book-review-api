package rest

import (
	"app/usecase"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(au usecase.AuthUseCase) AuthHandler {
	return &authHandler{
		authUseCase: au,
	}
}

func (ah authHandler) Login(c *gin.Context) {

	type RequestDataField struct {
		Mail     string `json:"mail" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var json RequestDataField
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	mail := json.Mail
	password := json.Password

	user, err := ah.authUseCase.Login(mail)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user_id": user.UserId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error!!! Not SECRET_KEY"})
		return
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not generate token"})
		return
	}

	user.Password = json.Password

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user, "userId": user.UserId})
}
