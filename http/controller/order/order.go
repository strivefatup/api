package order

import (
	"fmt"
	"strconv"
	"api/config"
	"encoding/json"
	"api/serveres/order"
	"api/tools/http/response"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Name 		string `form:"name" json:"name" xml:"name"  binding:"required"`
	Price 		float64 `form:"price" json:"price" xml:"price"  binding:"required"`
}


func GetOrderList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize, _ = config.Config("pagination.page_size").(int)
	}
	result, count := order.List(page, pageSize)
	c.JSON(200, response.ResponseList{result, response.Pagination{count, page, pageSize}})
}

func GetOrderById(c *gin.Context) {
	// fmt.Printf("%d", c.Param("id"))
	// id, _ := strconv.Atoi(c.Param("id"))
	// if data, err := order.Show(uint(id)); err != nil {
	// 	c.JSON(404, gin.H{"message": "record not found"})
	// } else {
	// 	c.JSON(200, data)
	// }
}

func CreateOrder(c *gin.Context) {
	var param Order
	c.Bind(&param)
	data, _ := json.Marshal(param)
	result := order.Create(data)
	c.JSON(201, result)
}

func UpdateOrder(c *gin.Context) {
	fmt.Printf("%d", c.Param("id"))
	// id, _ := strconv.Atoi(c.Param("id"))
	// var param Order
	// c.Bind(&param)
	// data, _ := json.Marshal(param)
	// if result, err := order.Update(uint(id), data); err != nil {
	// 	c.JSON(404, gin.H{"message": "record not found"})
	// } else {
	// 	c.JSON(201, result)
	// }
}

func DeleteOrderById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order.Delete(uint(id))
	c.JSON(204, gin.H{"message": ""})
}



