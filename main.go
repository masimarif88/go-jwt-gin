package main

import (
	"jwt-gin/controllers"
	"jwt-gin/middlewares"
	"jwt-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":8080")

}
