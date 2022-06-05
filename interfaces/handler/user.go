package rest

import (
	"app/middleware"
	"app/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"fmt"
)

type UserHandler interface {
	UserCreate(c *gin.Context)
	UserDetail(c *gin.Context)
	UserUpdate(c *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) UserCreate(c *gin.Context) {

	type RequestDataField struct {
		Name       string `json:"name" binding:"required"`
		Mail       string `json:"mail" binding:"required"`
		Password   string `json:"password" binding:"required"`
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
	hash, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 12)

	name := json.Name
	mail := json.Mail
	password := string(hash)

	user_id, err := uh.userUseCase.UserCreate(name, mail, password)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"user_id": user_id})
}

func (uh userHandler) UserDetail(c *gin.Context) {

	user_id := c.Param("user_id")

	user, err := uh.userUseCase.UserDetail(user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (uh userHandler) UserUpdate(c *gin.Context) {

	type RequestDataField struct {
		Name       string `json:"name" binding:"required"`
		Mail       string `json:"mail" binding:"required"`
		Password   string `json:"password" binding:"required"`
	}
	var json RequestDataField

	err := c.BindJSON(&json)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 12)

	user_id := middleware.ClaimUserID
	name := json.Name
	mail := json.Mail
	password := string(hash)

	err = uh.userUseCase.UserUpdate(user_id, name, mail, password)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "user update is success."})
}
