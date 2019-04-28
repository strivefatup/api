package user

import (
	"fmt"
	"strconv"
	"api/config"
	"encoding/json"
	"api/serveres/user"
	"api/tools/http/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `form:"user_name" json:"user_name" xml:"user_name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
}


func GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize, _ = config.Config("pagination.page_size").(int)
	}
	result, count := user.List(page, pageSize)
	c.JSON(200, response.ResponseList{result, response.Pagination{count, page, pageSize}})
}

func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if data, err := user.Show(uint(id)); err != nil {
		c.JSON(404, gin.H{"message": "record not found"})
	} else {
		c.JSON(200, data)
	}
}

func CreateUser(c *gin.Context) {
	var param User
	c.Bind(&param)
	data, _ := json.Marshal(param)
	result := user.Create(data)
	c.JSON(201, result)
}

func UpdateUser(c *gin.Context) {
	fmt.Printf("%d", c.Param("id"))
	// id, _ := strconv.Atoi(c.Param("id"))
	// var param User
	// c.Bind(&param)
	// data, _ := json.Marshal(param)
	// if result, err := user.Update(uint(id), data); err != nil {
	// 	c.JSON(404, gin.H{"message": "record not found"})
	// } else {
	// 	c.JSON(201, result)
	// }
}

func DeleteUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user.Delete(uint(id))
	c.JSON(204, gin.H{"message": ""})
}
