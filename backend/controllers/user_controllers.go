package controllers

import (
	"net/http"

	"care-vault/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Get user"})
}

func CreateUser(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

