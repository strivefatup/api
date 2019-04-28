package auth

import (
	"api/config"
	"encoding/json"
	"api/tools/crypto"
	"api/tools/jwt"
	userServer "api/serveres/user"
	userStruct "api/http/controller/user"
	
)

func Login(data []byte) (string, error) {
	user := &userStruct.User{}
	json.Unmarshal(data, user)
	user.Password = crypto.Md5(config.Config("jwt.salt").(string), user.Password)
	userModel, _ := userServer.GetUserByUserNameAndPwd(user.UserName, user.Password)
	token, err := JWT.GenerateToken(userModel.ID, userModel.UserName, userModel.Email)
	return token, err
}


func Registere(data []byte) (string, error) {
	user := &userStruct.User{}
	json.Unmarshal(data, user)
	user.Password = crypto.Md5(config.Config("jwt.salt").(string), user.Password)
	insertData, _ := json.Marshal(user)
	result := userServer.Create(insertData)
	token, err := JWT.GenerateToken(result.ID, result.UserName, result.Email)
	return token, err
}

func GetUserInfo(token string) (interface{}, error) {
	customClaims := JWT.ParseToken(token)
	return userServer.Show(customClaims.ID)
}