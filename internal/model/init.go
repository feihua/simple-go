package model

import (
	"fmt"
	"github.com/feihua/simple-go/pkg/config"
	"github.com/feihua/simple-go/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func Init() *gorm.DB {
	mysqlInfo := config.GlobalAppConfig.Mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlInfo.Username,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.Database,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		utils.Logger.Debugf("mysql连接异常%s", err.Error())
		panic(err)
	}

	utils.Logger.Debugf("mysql已连接: %s", dsn)

	return DB
	// 数据库迁移
	// migrate()
}

// func migrate() {
//	DB.AutoMigrate(&User{})
// }

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	utils.Logger.Debugf(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
