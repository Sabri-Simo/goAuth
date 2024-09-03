package controller

import (
	"github.com/gin-gonic/gin"
	"goAuth/initializers"
	"goAuth/models"
	"goAuth/service"
	"net/http"
)

func Singup(c *gin.Context) {
	var input struct {
		UserName        string `json:"user_name" binding:"required"`
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required,min=9"`
		PasswordConfirm string `json:"password_confirm" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}
	hashPassword, err := service.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if !service.IsvalidEmail(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	user := models.User{
		UserName: input.UserName,
		Email:    input.Email,
		Password: hashPassword,
	}

	var existingUser models.User
	if err := initializers.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
