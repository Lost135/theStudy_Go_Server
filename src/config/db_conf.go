package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB

func DBClient() {
	//连接数据库
	err := recover()
	Db, err = gorm.Open(mysql.Open(Yml.Mysql), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		NowFunc: func() time.Time {
			return time.Now()
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	CatchFatal(err)
	sqlDB, err := Db.DB()
	CatchErr(err)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(20)

	// SetConnMaxLifetime 设置了连接最大生存时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
