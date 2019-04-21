package main

import (
	"api/models"
	"api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化DB
	models.InitDB()

	//初始化路由
	router := gin.Default()
	routes.RegisterRoute(router)
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}




