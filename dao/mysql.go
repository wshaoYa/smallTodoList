package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// InitMySQL 连接数据库
func InitMySQL() {
	var (
		username = "root"
		pwd      = "sw6813329"
		db       = "person_practice"
		ip       = "localhost"
		port     = 3306
		err      error
	)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		username, pwd, ip, port, db)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
