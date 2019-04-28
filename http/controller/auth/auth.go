package auth

import (
	"encoding/json"
	"api/serveres/auth"
	"github.com/gin-gonic/gin"
	"api/http/controller/user"
)

type Token struct {
	Type        string `form:"type" json:"type" xml:"type"`
	AccessToken string `form:"access_token" json:"access_token" xml:"access_token"`
}


func Registere(c *gin.Context) {
	var param user.User
	c.Bind(&param)
	data, _ := json.Marshal(param)
	token, _ := auth.Registere(data)
	c.JSON(201, &Token{"Bearer", token})
}

func Login(c *gin.Context) {
	var param user.User
	c.Bind(&param)
	data, _ := json.Marshal(param)
	token, _ := auth.Login(data)
	c.JSON(201, &Token{"Bearer", token})
}

func GetUserInfo(c *gin.Context) {
	data, _ := auth.GetUserInfo(c.GetString("token"))
	c.JSON(200, data)
}