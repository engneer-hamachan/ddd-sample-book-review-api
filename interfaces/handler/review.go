package rest

import (
	"app/middleware"
	"app/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ReviewHandler interface {
	ReviewCreate(c *gin.Context)
	ReviewDetail(c *gin.Context)
	ReviewAll(c *gin.Context)
	ChangeReviewPublicFlg(c *gin.Context)
	ReviewDelete(c *gin.Context)
	CommentCreate(c *gin.Context)
	CommentDelete(c *gin.Context)
	ReviewLikeCreate(c *gin.Context)
	ReviewLikeDelete(c *gin.Context)
	CommentLikeCreate(c *gin.Context)
	CommentLikeDelete(c *gin.Context)
}

type reviewHandler struct {
	reviewUseCase usecase.ReviewUseCase
}

func NewReviewHandler(ru usecase.ReviewUseCase) ReviewHandler {
	return &reviewHandler{
		reviewUseCase: ru,
	}
}

func (rh reviewHandler) ReviewCreate(c *gin.Context) {

	type RequestDataField struct {
		BookTitle   string `json:"book_title" binding:"required"`
		ReviewTitle string `json:"review_title" binding:"required"`
		Publisher   string `json:"publisher" binding:"required"`
		Review      string `json:"review" binding:"required"`
		ReadedAt    string `json:"readed_at" binding:"required"`
		Stars       int    `json:"stars" binding:"required"`
		PublicFlg   bool `json:"public_flg" binding:"required"`
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

	current_user_id := middleware.ClaimUserID
	book_title := json.BookTitle
	review_title := json.ReviewTitle
	publisher := json.Publisher
	reviewVal := json.Review
	readed_at := json.ReadedAt
	stars := json.Stars
	public_flg := json.PublicFlg

	layout := "2006-01-02 15:04:05"
	convert_readed_at, err := time.Parse(layout, readed_at)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	review_id, err := rh.reviewUseCase.ReviewCreate(current_user_id, book_title, review_title, publisher, reviewVal, convert_readed_at, stars, public_flg)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"review_id": review_id})
}

func (rh reviewHandler) ReviewDetail(c *gin.Context) {

	review_id := c.Param("review_id")
	current_page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	review, err := rh.reviewUseCase.ReviewDetail(review_id, current_page)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, review)
}

func (rh reviewHandler) ReviewAll(c *gin.Context) {

	current_page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	reviews, err := rh.reviewUseCase.ReviewAll(current_page)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, reviews)
}

func (rh reviewHandler) ChangeReviewPublicFlg(c *gin.Context) {

	type RequestDataField struct {
		ReviewId string `json:"review_id" binding:"required"`
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

	review_id := json.ReviewId

	err = rh.reviewUseCase.ChangeReviewPublicFlg(review_id, middleware.ClaimUserID)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "change public flg is success."})
}

func (rh reviewHandler) ReviewDelete(c *gin.Context) {

	type RequestDataField struct {
		ReviewId string `json:"review_id" binding:"required"`
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

	current_user_id := middleware.ClaimUserID
	review_id := json.ReviewId

	err = rh.reviewUseCase.ReviewDelete(review_id, current_user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "delete review is success."})
}

func (rh reviewHandler) CommentCreate(c *gin.Context) {

	type RequestDataField struct {
		ReviewId   string `json:"review_id" binding:"required"`
		Comment string `json:"comment" binding:"required"`
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

	review_id := json.ReviewId
	current_user_id := middleware.ClaimUserID
	comment := json.Comment

	comment_id, err := rh.reviewUseCase.CommentCreate(review_id, current_user_id, comment)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"comment_id": comment_id})
}

func (rh reviewHandler) CommentDelete(c *gin.Context) {

	type RequestDataField struct {
		CommentId string `json:"comment_id" binding:"required"`
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

	current_user_id := middleware.ClaimUserID
	comment_id := json.CommentId

	err = rh.reviewUseCase.CommentDelete(comment_id, current_user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "delete comment is success."})
}

func (rh reviewHandler) ReviewLikeCreate(c *gin.Context) {

	type RequestDataField struct {
		ReviewId   string `json:"review_id" binding:"required"`
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

	review_id := json.ReviewId
	current_user_id := middleware.ClaimUserID

	err = rh.reviewUseCase.ReviewLikeCreate(review_id, current_user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "creeate review like is success"})
}

func (rh reviewHandler) ReviewLikeDelete(c *gin.Context) {

	type RequestDataField struct {
		ReviewId   string `json:"review_id" binding:"required"`
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

	review_id := json.ReviewId
	current_user_id := middleware.ClaimUserID

	err = rh.reviewUseCase.ReviewLikeDelete(review_id, current_user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "delete review like is success"})
}

func (rh reviewHandler) CommentLikeCreate(c *gin.Context) {

	type RequestDataField struct {
		CommentId   string `json:"comment_id" binding:"required"`
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

	comment_id := json.CommentId
	current_user_id := middleware.ClaimUserID

	err = rh.reviewUseCase.CommentLikeCreate(comment_id, current_user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "creeate comment like is success"})
}

func (rh reviewHandler) CommentLikeDelete(c *gin.Context) {

	type RequestDataField struct {
		CommentId   string `json:"comment_id" binding:"required"`
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

	comment_id := json.CommentId
	current_user_id := middleware.ClaimUserID

	err = rh.reviewUseCase.CommentLikeDelete(comment_id, current_user_id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "delete comment like is success"})
}
