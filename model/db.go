package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error
var dsn = utils.DbUser + ":" +
	utils.DbPassWord + "@tcp(" +
	utils.DbHost +
	utils.DbPort + ")/" +
	utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

func InitDb() {
	//dsn = "root:123456@tcp(127.0.0.1:3306)/ginblog?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//fmt.Println(dsn)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用默认表名的复数形式
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
		return
	}

	// 自动迁移
	err := db.AutoMigrate(&User{}, &Category{}, &Article{}, &Profile{})
	if err != nil {
		fmt.Println("迁移失败！", err)
		return
	}

	sqlDB, ERR := db.DB()

	if ERR != nil {
		fmt.Println("数据库池错误，请检查", ERR)
		return
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(8 * time.Hour)

	//fmt.Println(db)

}
