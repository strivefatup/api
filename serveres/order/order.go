package order

import (
	"api/models/order"
)

var orderModel order.OrderModel

func Show(id uint) (interface{}, error) {
	return orderModel.Find(id)
}

func List(page, pageSize int) (interface{}, int) {
	return orderModel.GetList(page, pageSize)
}
func Create(order []byte) interface{} {
	return orderModel.Create(order)
}

func Update(id uint, order []byte) (interface{}, error) {
	return orderModel.Update(id, order)
}

func Delete(id uint) {
	orderModel.Delete(id)
}