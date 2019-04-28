package user

import (
	"api/models/user"
)

var userModel user.UserModel

func Show(id uint) (interface{}, error) {
	return userModel.Find(id)
}

func List(page, pageSize int) (interface{}, int) {
	return userModel.GetList(page, pageSize)
}
func Create(user []byte) *user.UserModel {
	return userModel.Create(user)
}

func Update(id uint, user []byte) (interface{}, error) {
	return userModel.Update(id, user)
}

func Delete(id uint) {
	userModel.Delete(id)
}

func GetUserByUserNameAndPwd(userName, password string) (*user.UserModel, error){
	return userModel.GetUserByUserNameAndPwd(userName, password)

}