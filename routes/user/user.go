package user

import (
	"github.com/gin-gonic/gin"
	"api/http/controller/user"
	"api/http/controller/auth"
)

func User(engine *gin.Engine){
	route := engine.Group("/api")
    {
        route.GET("/token/user", auth.GetUserInfo)
        route.GET("/user", user.GetUserList)
        route.GET("/user/:id", user.GetUserById)
		route.POST("/user", user.CreateUser)
        route.PUT("/user/:id", user.UpdateUser)
        route.DELETE("/user/:id", user.DeleteUserById)
    }
}