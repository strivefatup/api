package database

import "github.com/jinzhu/gorm"

type DbconnPool struct {
}

type ConnPool interface {
	GetInstance() *DbconnPool
	InitDataPool() bool
	GetDB() *gorm.DB
}