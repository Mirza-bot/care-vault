package controllers

import (
	"log"
	"net/http"
	"strconv"

	userdtos "care-vault/dtos/user_dtos"
	"care-vault/models"
	"care-vault/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetUser(c *gin.Context) {
    var userId = c.Param("id")
    log.Printf("the userId is %v", userId)
    id, err := strconv.Atoi(userId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var user models.User

    if err := db.First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Databse query failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}


func CreateUser(c *gin.Context) {
    var userDto userdtos.UserCreateDto

    if err := c.ShouldBindJSON(&userDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if userDto.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "A name is required"})
    }

    user := models.User{
        Name: userDto.Name,
        Email: userDto.Email,
        // TODO: Create a HashPassword-function to hash the password here
        Password: userDto.Password,
    }

    // TODO: Make sure to log properly if the user could not be created because the email is already used
    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}


func DeleteUser(c *gin.Context) {
    userID := c.Param("id")
    id, err := strconv.Atoi(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := db.Delete(&models.User{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete User"})
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


func ModifyUser(c *gin.Context) {
    var userDto userdtos.UserModifyDto
    if err := c.ShouldBindJSON(&userDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.Param("id")
    id, err := strconv.Atoi(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var user models.User
    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if userDto.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "A name is required"})
        return
    }

    if userDto.Email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "E-Mail is required"})
        return
    }

    if utils.IsValidEmail(userDto.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email adress"})
        return
    }

    user.Name = userDto.Name
    user.Email = userDto.Email

    if err := db.Model(&user).Updates(user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

