package routes

import (
	"care-vault/controllers"
	"log"
    _ "care-vault/docs"

	"github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func SetupAuthRoutes(router *gin.Engine) {
    log.Println("Registering login route")
    auth := router.Group("/auth")
    {
        auth.GET("/login", controllers.Login)
        auth.GET("/register", controllers.Register)
    }
}

func SetupUserRoutes(router *gin.Engine) {
	log.Println("Registering user routes")
    user := router.Group("/user")
    {
        user.POST("/", controllers.CreateUser)
        user.GET("/:id", controllers.GetUser)
        user.DELETE("/:id", controllers.DeleteUser)
        user.PATCH("/:id", controllers.ModifyUser)
    }
}

func SetupTaskRoutes(router *gin.Engine) {
    log.Println("Registering task routes")
    // task := router.Group("/task")
    {
        // task.POST("/", controllers.CreateTask)
        // task.GET("/:id", controllers.GetTask)
        // task.DELETE("/:id", controllers.DeleteTask)
        // task.PATCH("/:id", controllers.ModifyTask)
    }
}

func SetupCompanyRoutes(router *gin.Engine) {
    log.Println("Registering company routes")
    // company := router.Group("/company")
    {
        // General
        // company.POST("/", controllers.RegisterCompany)
        // company.GET("/:id", controllers.GetCompany)
        // company.DELETE("/:id", controllers.DeleteCompany)
        // company.PATCH("/:id", controllers.ModifyCompany)

        // Employee
        // company.POST("/employee", controllers.CreateEmployee)
        // company.GET("/employee/:id", controllers.GetEmployee)
        // company.DELETE("/employee/:id", controllers.RemoveEmployee)
        // company.PATCH("/eomployee/:id", controllers.ModifyEmployee)
    }
}


func SetupRoutes(router *gin.Engine) {
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Setting up routes...")
    SetupAuthRoutes(router)
    SetupUserRoutes(router)
    SetupTaskRoutes(router)
    SetupCompanyRoutes(router)
    log.Println("Routes setup completed")
}
