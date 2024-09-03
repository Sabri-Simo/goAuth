package controller

import (
	"github.com/gin-gonic/gin"
	"goAuth/initializers"
	"goAuth/models"
	"goAuth/service"
	"net/http"
)

func Login(c *gin.Context) {
	var input struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required,min=9"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := initializers.DB.Where("user_name=?", input.UserName).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	if !service.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "you log"})
}
