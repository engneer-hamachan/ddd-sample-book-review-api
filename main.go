package main

import (
	"app/config"
	"app/infrastructures/domain_service"
	"app/infrastructures/query_service"
	"app/infrastructures/repository"
	handler "app/interfaces/handler"
	"app/middleware"
	"app/usecase"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	//DB
	db := config.Connect()
	defer db.Close()
	//DI
	userRepository := repository.NewUserRepository(db)
	userDomainService := domain_service.NewUserDomainService(db)
	userQueryService := query_service.NewUserQueryService(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userDomainService, userQueryService)
	userHandler := handler.NewUserHandler(userUseCase)

	reviewRepository := repository.NewReviewRepository(db)
	reviewDomainService := domain_service.NewReviewDomainService(db)
	reviewQueryService := query_service.NewReviewQueryService(db)
	reviewUseCase := usecase.NewReviewUseCase(reviewRepository, reviewDomainService, reviewQueryService)
	reviewHandler := handler.NewReviewHandler(reviewUseCase)

	authQueryService := query_service.NewAuthQueryService(db)
	authUseCase := usecase.NewAuthUseCase(authQueryService)
	authHandler := handler.NewAuthHandler(authUseCase)

	router.POST("/login", authHandler.Login)
	router.POST("/users", userHandler.UserCreate)

	authUserGroup := router.Group("/")
	authUserGroup.Use(middleware.LoginCheckMiddleware())
	{
		authUserGroup.GET("/user/:user_id", userHandler.UserDetail)
		authUserGroup.PUT("/user", userHandler.UserUpdate)

		authUserGroup.POST("/review", reviewHandler.ReviewCreate)
		authUserGroup.GET("/review/:review_id/page/:page", reviewHandler.ReviewDetail)
		authUserGroup.GET("/reviews/page/:page", reviewHandler.ReviewAll)
		authUserGroup.PUT("/review/public_flg", reviewHandler.ChangeReviewPublicFlg)
		authUserGroup.DELETE("/review", reviewHandler.ReviewDelete)
		authUserGroup.POST("/review/like", reviewHandler.ReviewLikeCreate)
		authUserGroup.DELETE("/review/like", reviewHandler.ReviewLikeDelete)

		authUserGroup.POST("/review/comment", reviewHandler.CommentCreate)
		authUserGroup.DELETE("/review/comment", reviewHandler.CommentDelete)
		authUserGroup.POST("/review/comment/like", reviewHandler.CommentLikeCreate)
		authUserGroup.DELETE("/review/comment/like", reviewHandler.CommentLikeDelete)
	}

	router.Run("localhost:8080")
	fmt.Println("start")
}
