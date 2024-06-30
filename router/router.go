package router

import (
	"belajar/golang/controller"
	"belajar/golang/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:userId", middleware.UserAuthorization(), controller.UpdateUser)
		userRouter.DELETE("/:userId", middleware.UserAuthorization(), controller.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetPhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.GET("/", controller.GetComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controller.DeleteComent)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.GET("/", controller.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}

	return r
}
