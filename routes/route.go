package routes

import (
	"github.com/gin-gonic/gin"
	"api/routes/order"
)

func RegisterRoute(engine *gin.Engine) *gin.Engine {
	order.Order(engine)
	return engine
}


