package controllers

import (
	"net/http"
	"strconv"
	"time"

	userdtos "care-vault/dtos/user_dtos"
	"care-vault/models"
	"care-vault/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


// @Summary Get a user by ID
// @Description Get detailed information about a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "User ID"
// @Success 200 {object} models.User "Successfully retrieved user"
// @Failure 400 "Invalid ID format"
// @Failure 404 "User not found"
// @Failure 404 "Database query failed"
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
    var userId = c.Param("id")
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



// @Summary Create a new user
// @Description Creates a new user with the provided name and email. Ensures the email is unique.
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user  body  userdtos.UserPublicDto  true  "User data"
// @Success 200 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]string "Invalid input or email already in use"
// @Failure 500 {object} map[string]string "Could not create user"
// @Router /user [post]
func CreateUser(c *gin.Context) {
    var userDto userdtos.UserPublicDto

    if err := c.ShouldBindJSON(&userDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var exsistingUser models.User
    if err := db.Where("email = ?", userDto.Email).First(&exsistingUser).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "E-Mail is already in use."})
        return
    }

    if userDto.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "A name is required"})
        return
    }

    user := models.User{
        Name: userDto.Name,
        Email: userDto.Email,
        Created: time.Now(),
    }

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

    if !utils.IsValidEmail(userDto.Email) {
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

