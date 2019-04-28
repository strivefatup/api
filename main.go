package main

import (
	"api/models"
	"api/routes"
	"api/http/middleware"
	"api/http/controller/auth"
	"github.com/gin-gonic/gin"
)

func main() {

	//初始化DB
	models.InitDB()

	//初始化路由
	router := gin.Default()
	router.POST("api/registere", auth.Registere)
	router.POST("api/login", auth.Login)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "404 Not Found"})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "404 Not Found"})
	})
	router.Use(middleware.Auth())
	routes.RegisterRoute(router)
	//Listen and Server in 0.0.0.0:8085
	router.Run(":8085")
}




