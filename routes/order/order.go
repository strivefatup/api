package order

import (
	"github.com/gin-gonic/gin"
	"api/http/controller/order"
)

func Order(engine *gin.Engine){
	route := engine.Group("/api")
    {
        route.GET("/order", order.GetOrderList)
        route.GET("/order/:id", order.GetOrderById)
		route.POST("/order", order.CreateOrder)
        route.PUT("/order/:id", order.UpdateOrder)
        route.DELETE("/order/:id", order.DeleteOrderById)
    }
}