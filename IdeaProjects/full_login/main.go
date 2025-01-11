package main

import (
	"full_login/controllers"
	"full_login/middlewares"
	"full_login/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	protected.POST("/auth", controllers.AuthToken)

	r.Run(":8080")

}
