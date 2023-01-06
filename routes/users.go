package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mmrumii/go-gin/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.POST("/user/add", controller.AddUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
}
