package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

// DB 后续用于操作数据库
var DB *gorm.DB

func InitDB(dsn string) {
	//DB = gohelper_db.InitDB(dsn)
	// 连接数据库
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 打印 SQL
	db = db.Debug()
	// 设置连接池大小
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get DB instance")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 设置WAL模式
	db.Exec("PRAGMA journal_mode=WAL")
	//_ = db.Exec("PRAGMA journal_size_limit=104857600;")
	//_ = db.Exec("PRAGMA busy_timeout=999999;")
	// 设置事务隔离级别
	err = db.Exec("PRAGMA read_uncommitted = 1").Error
	if err != nil {
		panic("failed to set transaction isolation level")
	}

	DB = db
	//// 创建表
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic("failed to create table")
	//}
	//
	//// 插入数据
	//err = db.Transaction(func(tx *gorm.DB) error {
	//	user := User{Name: "John"}
	//	err := tx.Create(&user).Error
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//if err != nil {
	//	panic("failed to insert data")
	//}
}
