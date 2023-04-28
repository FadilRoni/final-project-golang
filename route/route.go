package route

import (
	"final-project/handler"
	"final-project/middleware"

	"github.com/gin-gonic/gin"

	_ "final-project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Book Example API
// @version         1.0
// @description     This is a simple service for managing book.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	// server := handler.NewHttpServer(app)

	users := r.Group("/users") // prefix
	{
		users.POST("/register", server.Register) // /register
		users.POST("/login", server.Login)       // /login
	}

	smedia := r.Group("/social_media") // prefix
	{
		smedia.Use(middleware.Authentication())
		smedia.POST("", server.CreateSocialMedia)                                              // /social_media
		smedia.GET("", server.GetAllSocialMedia)                                               // /social_media
		smedia.GET("/:id", middleware.SocialMediaAuthorization(), server.GetOneSocialMedia)    // /social_media/:id
		smedia.PUT("/:id", middleware.SocialMediaAuthorization(), server.UpdateSocialMedia)    // /social_media/:id
		smedia.DELETE("/:id", middleware.SocialMediaAuthorization(), server.DeleteSocialMedia) // /social_media/:id

	}

	photo := r.Group("/photo") // prefix
	{
		photo.Use(middleware.Authentication())
		photo.POST("", server.CreatePhoto)                                        // /photo
		photo.GET("", server.GetAllPhoto)                                         // /photo
		photo.GET("/:id", middleware.PhotoAuthorization(), server.GetOnePhoto)    // /photo/:id
		photo.PUT("/:id", middleware.PhotoAuthorization(), server.UpdatePhoto)    // /photo/:id
		photo.DELETE("/:id", middleware.PhotoAuthorization(), server.DeletePhoto) // /photo/:id

	}

	comment := r.Group("/comment") // prefix
	{
		comment.Use(middleware.Authentication())
		comment.POST("/:idPhoto", server.CreateComment)                                                 // /comment/:idPhoto
		comment.GET("/:idPhoto", server.GetAllComment)                                                  // /comment/:idPhoto
		comment.GET("/:idPhoto/:idComment", middleware.CommentAuthorization(), server.GetOneComment)    // /comment/:idPhoto/:idComment
		comment.PUT("/:idPhoto/:idComment", middleware.CommentAuthorization(), server.UpdateComment)    // /comment/:idPhoto/:idComment
		comment.DELETE("/:idPhoto/:idComment", middleware.CommentAuthorization(), server.DeleteComment) // /comment/:idPhoto/:idComment

	}

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
