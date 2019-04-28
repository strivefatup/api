package routes

import (
	"api/routes/user"
	"api/routes/order"
	"github.com/gin-gonic/gin"
	
)

func RegisterRoute(engine *gin.Engine) *gin.Engine {
	order.Order(engine)
	user.User(engine)
	return engine
}


