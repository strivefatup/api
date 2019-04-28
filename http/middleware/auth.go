package middleware

import (
	// "fmt"
	"reflect"
	"strings"
	"api/tools/jwt"
	"github.com/gin-gonic/gin"
)



func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization, ok := c.Request.Header["Authorization"]
		if ok {
			token := reflect.ValueOf(authorization[0]).Interface().(string)
			parseToken := reflect.ValueOf(strings.SplitN(token, " ", 2)[1]).Interface().(string)
			if ! JWT.Check(parseToken) {
				c.JSON(401, gin.H{"message": "Token invalid"})
				c.Abort()
			} else {
				c.Set("token", parseToken)
				c.Next()
			}
		} else {
			c.JSON(401, gin.H{"message": "Token invalid"})
			c.Abort()
		}
	}
}
