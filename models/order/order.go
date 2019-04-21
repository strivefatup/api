package order

import (
	"time"
	"api/models"
	"encoding/json"
)


// 定义原始的数据库字段
type OrderModel struct {
	models.BaseModel
	Name string 		`json:"name";gorm:"type:varchar(32);NOT NULL;column:name"`
	Price float64       `json:"price";gorm:"type:decimal(64);NOT NULL;DEFAULT:0;column:price"`
}

/**
*@des 设置当前模型的表名
*/
func (order *OrderModel) TableName() string {
	return "order"
}


/**
*@des 根据id查找对应的记录
*/
func (order *OrderModel) Find(id uint) (*OrderModel, error) {
	result := &OrderModel{}
	err := order.GetMysqlDb().Where("id = ?", id).First(&result).Error
	return result, err
}

/**
*@des 根据页码和每页记录数返回数据
*/
func (order *OrderModel) GetList(page, pageSize int) ([]*OrderModel, int) {
	var count int
	result := []*OrderModel{}
	order.GetMysqlDb().Model(&order).Count(&count)
	order.GetMysqlDb().Limit(pageSize).Offset((page - 1) * pageSize).Find(&result)
	return result, count
}

/**
*@des 创建一条新的记录
*/
func (order *OrderModel) Create(data []byte) *OrderModel {
	result := &OrderModel{}
	json.Unmarshal(data, result)
	order.GetMysqlDb().Create(&result)
	return result
}

/**
*@des 根据id修改对应的记录
*/
func (order *OrderModel) Update(id uint, data []byte) (*OrderModel, error) {
	order.ID = id
	result := &OrderModel{}
	json.Unmarshal(data, result)
	rowsAffected := order.GetMysqlDb().Model(&order).UpdateColumns(&result).RowsAffected
	if rowsAffected == 1 {
		order.GetMysqlDb().Model(&order).UpdateColumn("UpdatedAt", time.Now().Unix())
	}
	return order.Find(id)
}


/**
*@des 根据id删除对应的记录(软删除)
*/
func (order *OrderModel) Delete(id uint) {
	order.ID = id
	order.GetMysqlDb().Model(&order).UpdateColumn("DeletedAt", time.Now().Unix())
}
