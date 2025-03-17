package routes

import (
	"gin-lecture/internal/delivery"
	"gin-lecture/internal/repository"
	"gin-lecture/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	var postRepo service.PostRepository = repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := delivery.NewPostHandler(postService)

	posts := r.Group("api/v1/posts")
	{
		posts.GET("/", postHandler.GetPosts)
		posts.GET("/:id", postHandler.GetPost)
		posts.POST("/", postHandler.CreatePost)
		posts.PUT("/:id", postHandler.UpdatePost)
		posts.DELETE("/:id", postHandler.DeletePost)
	}
}
