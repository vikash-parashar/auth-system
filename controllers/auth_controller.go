// controllers/auth_controller.go
package controllers

import (
	"auth-system/models"
	"auth-system/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Register user with email confirmation
func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassqord, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatalln(err)
		return
	}
	user.Password = hashedPassqord
	DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

// controllers/auth_controller.go
func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve user from the database by username
	DB.Where("email = ?", user.Email).First(&user)

	// Verify the password
	if err := utils.VerifyPassword(user.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token and send it as a response
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "token": token})
}
func UpdateUser(c *gin.Context) {
	//TODO:
}
func DeleteUser(c *gin.Context) {
	//TODO:
}
