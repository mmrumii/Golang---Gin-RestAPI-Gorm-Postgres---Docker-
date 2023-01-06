package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mmrumii/go-gin/config"
	"github.com/mmrumii/go-gin/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func AddUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(201, "User Created Successfully.")
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, "User Updated Successfully....")
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ? ", c.Param("id")).Delete(&user)
	c.JSON(200, "User Deleted Successfully.")
}
