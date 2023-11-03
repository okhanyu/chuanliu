package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// DB 后续用于操作数据库
var DB *gorm.DB

func InitDB(dsn string) {
	initMySQL(dsn)
}

func initSQLite(dsn string) {
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

func initMySQL(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接到数据库")
	}
	db = db.Exec("SET GLOBAL character_set_server= 'utf8mb4'")
	db = db.Exec("SET GLOBAL character_set_database = 'utf8mb4'")
	if err != nil {
		fmt.Println("修改character_set_system设置失败:", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		// 处理错误
	}
	// 打印 SQL
	db = db.Debug()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetConnMaxLifetime(time.Hour)

	// 定义定时任务
	go func() {
		for {
			// 执行 SELECT 1 查询以保持连接活跃
			db.Exec("SELECT 1")

			// 等待一段时间后再次执行
			time.Sleep(10 * time.Minute)
		}
	}()

	DB = db
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	panic("无法迁移数据库")
	//}
}
