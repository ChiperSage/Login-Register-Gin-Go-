package main

import (
    "github.com/gin-gonic/gin"
    "gin-user-management/config"
    "gin-user-management/controllers"
    "gin-user-management/middlewares"
)

func main() {
    r := gin.Default()

    config.ConnectDatabase()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    authorized := r.Group("/")
    authorized.Use(middlewares.Auth())
    {
        authorized.GET("/users", controllers.GetUsers)
        authorized.GET("/users/:id", controllers.GetUser)
        authorized.PUT("/users/:id", controllers.UpdateUser)
        authorized.DELETE("/users/:id", controllers.DeleteUser)
    }

    r.Run()
}
