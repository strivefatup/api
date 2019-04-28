package middleware

import (
    "fmt"
    "strconv"
	"github.com/gin-gonic/gin"
)

func ConvertIdStringsToUint() gin.HandlerFunc {
	return func(c *gin.Context) {
        fmt.Printf("%d", 222)
        id, _ := strconv.Atoi(c.Param("id"))
        // c.Param("id") = id
        c.Set("id", uint(id))
        fmt.Printf("%c", c)
	    c.Next()
	}
}