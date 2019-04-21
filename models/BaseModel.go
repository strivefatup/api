package models

import (
	// "fmt"
	"time"
	"api/config"
	"database/sql"
	"github.com/jinzhu/gorm"
	"api/tools/databases/mysql"
)


var timestamp int64 = time.Now().Unix()


type BaseModel struct {
	ID        uint  `json:"id";gorm:"_;primary_key"`
	CreatedAt int64 `json:"created_at";gorm:"type:int(10);column:created_at"`
	UpdatedAt int64 `json:"updated_at";gorm:"type:int(10);column:updated_at"`
	DeletedAt sql.NullInt64 `json:"deleted_at";gorm:"type:int(10);column:deleted_at"`
}


func InitDB() {
	var logger *config.Logger
	isSuccess := mysql.GetInstance().InitDataPool()
	if !isSuccess{
		logger.Debug("init database pool failure...")
	}
}

func (model *BaseModel) GetMysqlDb() (db *gorm.DB) {
	return mysql.GetInstance().GetDB()
}

func (model *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", timestamp)
	scope.SetColumn("UpdatedAt", timestamp)
	return nil
}









