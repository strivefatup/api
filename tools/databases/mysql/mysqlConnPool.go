package mysql

import (
	"fmt"
	"sync"
	"api/config"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

/*
* MysqlConnPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlConnPool struct {
}

var db *gorm.DB
var db_err error
var once sync.Once
var instance *MysqlConnPool



func GetInstance() *MysqlConnPool {
    once.Do(func() {
        instance = &MysqlConnPool{}
    })
    return instance
}

/*
* @fuc 初始化数据库连接(可在mail()适当位置调用)
*/
func (connPool *MysqlConnPool) InitDataPool() (issucc bool) {
	var logger *config.Logger
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config("db.mysql.connection.username"), 
		config.Config("db.mysql.connection.password"), config.Config("db.mysql.connection.host"), 
		config.Config("db.mysql.connection.port"), config.Config("db.mysql.connection.dbname"),
	)
    db, db_err = gorm.Open("mysql", dns)
    if db_err != nil {
        logger.Debug(db_err)
        return false
    }
    //关闭数据库，db会被多个goroutine共享，可以不调用
    // defer db.Close()
    return true
}

/*
* @fuc  对外获取数据库连接对象db
*/
func (connPool *MysqlConnPool) GetDB() (db_con *gorm.DB) {
    return db
}
