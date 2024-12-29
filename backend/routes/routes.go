package routes

import (
	"care-vault/controllers"
    "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	log.Println("Setting up routes...")

    api := router.Group("/api")

    api.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    // User-AP
	log.Println("Registering user routes")
    {
        api.GET("/users/:id", controllers.GetUser)
        api.POST("/users", controllers.CreateUser)
        api.DELETE("/users/:id", controllers.DeleteUser)
        api.PATCH("/users/:id", controllers.ModifyUser)
    }


    log.Println("Routes setup completed")
}
