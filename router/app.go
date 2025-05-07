package router

import (
	"github.com/gin-gonic/gin"
	"myBlog/middleware"
	"myBlog/service"
)

func App() *gin.Engine {
	r := gin.Default()
	defaultGroup := r.Group("/api")
	defaultGroup.Use(middleware.JwtAuth())
	defaultGroup.POST("/addPost", service.AddPost)
	defaultGroup.GET("/getPost", service.GetPost)
	defaultGroup.PUT("/updatePost", service.UpdatePost)
	defaultGroup.DELETE("/delete/:id", service.DeletePost)
	defaultGroup.GET("/getCommentList", service.GetCommentList)
	defaultGroup.POST("/addComment", service.AddComment)

	noAuthGroup := r.Group("api")
	noAuthGroup.POST("/login", service.Login)
	noAuthGroup.POST("/register", service.Register)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	return r
}
