package user

import (
	"fmt"
	"time"
	"api/models"
	"encoding/json"
)


type UserModel struct {
	models.BaseModel
    UserName string `json:"user_name";gorm:"type:varchar(32);NOT NULL;column:user_name"`
	Password string `json:"password";gorm:"type:varchar(32);NOT NULL;column:password"`
	Email    string `json:"email";gorm:"type:varchar(32);NOT NULL;column:email"`
}

// type Transform struct {
// 	models.BaseModel
//     UserName string `json:"user_name";gorm:"type:varchar(32);NOT NULL;column:user_name"`
// 	Password string `json:"password";gorm:"type:varchar(32);NOT NULL;column:password"`
// 	Email    string `json:"email";gorm:"type:varchar(32);NOT NULL;column:email"`
// }


func (user *UserModel) TableName() string {
	return "user"
}


/**
*@des 根据id查找对应的记录
*/
func (user *UserModel) Find(id uint) (*UserModel, error) {
	result := &UserModel{}
	err := user.GetMysqlDb().Where("id = ?", id).First(&result).Error
	return result, err
}

/**
*@des 根据页码和每页记录数返回数据
*/
func (user *UserModel) GetList(page, pageSize int) ([]*UserModel, int) {
	var count int
	result := []*UserModel{}
	user.GetMysqlDb().Model(&user).Count(&count)
	user.GetMysqlDb().Limit(pageSize).Offset((page - 1) * pageSize).Find(&result)
	return result, count
}

/**
*@des 创建一条新的记录
*/
func (user *UserModel) Create(data []byte) *UserModel {
	result := &UserModel{}
	json.Unmarshal(data, result)
	fmt.Printf("%v", result)
	user.GetMysqlDb().Create(&result)
	return result
}

/**
*@des 根据id修改对应的记录
*/
func (user *UserModel) Update(id uint, data []byte) (*UserModel, error) {
	user.ID = id
	result := &UserModel{}
	json.Unmarshal(data, result)
	rowsAffected := user.GetMysqlDb().Model(&user).UpdateColumns(&result).RowsAffected
	if rowsAffected == 1 {
		user.GetMysqlDb().Model(&user).UpdateColumn("UpdatedAt", time.Now().Unix())
	}
	return user.Find(id)
}


/**
*@des 根据id删除对应的记录(软删除)
*/
func (user *UserModel) Delete(id uint) {
	user.ID = id
	user.GetMysqlDb().Model(&user).UpdateColumn("DeletedAt", time.Now().Unix())
}

func (user *UserModel) GetUserByUserNameAndPwd(userName, password string) (*UserModel, error) {
	result := &UserModel{}
	err := user.GetMysqlDb().Where("user_name = ? AND password = ?", userName, password).First(&result).Error
	return result, err
}

