package routes

import (
	"care-vault/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")

    // User-AP
    {
        api.GET("/users/:id", controllers.GetUser)
        api.POST("/users/create", controllers.CreateUser)
    }
    // api.POST("/users", controllers.CreateUser)
}
