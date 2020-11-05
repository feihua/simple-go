package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"simple-go/pkg/config"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Mysql.Username,
			config.Mysql.Password,
			config.Mysql.Host,
			config.Mysql.Port,
			config.Mysql.Database,
		),
	)
	if err != nil {
		panic(err)
	}

	//设置连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	//设置表名为单数
	DB.SingularTable(true)
	//开启sql日志
	DB.LogMode(true)

	//数据库迁移
	//migrate()
}

//func migrate() {
//	DB.AutoMigrate(&User{})
//}

func Close() {
	defer DB.Close()
}
