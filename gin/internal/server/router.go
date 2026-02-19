package server

import (
	"github.com/gin-gonic/gin"
	"note_api/internal/handlers"
	"note_api/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	blogRepo := repository.NewBlogRepository(db)
	blogHandler := handlers.NewBlogHandler(blogRepo)

	api := r.Group("/api/v1")
	{
		api.POST("/blogs", blogHandler.Create)
		api.GET("/blogs", blogHandler.GetAll)
		api.GET("/blogs/:id", blogHandler.GetByID)
		api.PUT("/blogs/:id", blogHandler.Update)
		api.DELETE("/blogs/:id", blogHandler.Delete)
	}

	return r
}
