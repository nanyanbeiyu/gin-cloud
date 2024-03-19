package model

import (
	"fmt"
	"go_cloud/util/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB_USER,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	// 自动迁移
	db.AutoMigrate(&User{})

	// 连接池
	sqlDb, _ := db.DB()
	// 设置连接池大小
	sqlDb.SetMaxIdleConns(10)
	// 设置最大连接数
	sqlDb.SetMaxOpenConns(100)
	// 设置连接最大空闲时间
	sqlDb.SetConnMaxLifetime(time.Hour)
}

// 取消复数表名
func (User) TableName() string {
	return "user"
}
