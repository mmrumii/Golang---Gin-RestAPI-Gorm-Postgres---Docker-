package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mmrumii/go-gin/controller"
)

func BlogRoute(router *gin.Engine) {
	router.GET("/blogs", controller.GetBlogs)
	router.POST("/blog/add", controller.AddBlog)
	router.PUT("/blog/:id", controller.UpdateBlog)
	router.DELETE("/blog/:id", controller.DeleteBlog)
}
