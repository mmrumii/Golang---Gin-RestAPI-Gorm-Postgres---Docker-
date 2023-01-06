package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mmrumii/go-gin/config"
	"github.com/mmrumii/go-gin/routes"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	config.Connect()
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
