package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mmrumii/go-gin/config"
	"github.com/mmrumii/go-gin/models"
)

func GetBlogs(c *gin.Context) {
	blogs := []models.Blog{}
	config.DB.Find(&blogs)
	c.JSON(200, &blogs)
}

func AddBlog(c *gin.Context) {
	blog := models.Blog{}
	c.BindJSON(&blog)
	config.DB.Create(&blog)
	c.JSON(201, "Blog Created Successfully.")
}

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	config.DB.Where("id = ?", c.Param("id")).First(&blog)
	c.BindJSON(&blog)
	config.DB.Save(&blog)
	c.JSON(200, "Blog Updated Successfully....")
}

func DeleteBlog(c *gin.Context) {
	var blog models.Blog
	config.DB.Where("id = ? ", c.Param("id")).Delete(&blog)
	c.JSON(200, "Blog Deleted Successfully.")
}
