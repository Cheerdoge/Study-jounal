package config

import (
	"fmt"
	"item-repository/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	// 1. 构建不带数据库名的 DSN，用于初始连接 MySQL 服务器
	dsnBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		APPConfig.DBUser, APPConfig.DBPassword, APPConfig.DBHost, APPConfig.DBPort,
	)

	// 连接到 MySQL 服务器
	dbBase, err := gorm.Open(mysql.Open(dsnBase))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mysql server: %w", err)
	}

	// 2. 执行创建数据库的 SQL 语句 (如果不存在则创建)
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;", APPConfig.DBName)
	if err := dbBase.Exec(createDBSQL).Error; err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}
	log.Printf("Ensured database '%s' exists", APPConfig.DBName)

	// 获取底层的 *sql.DB 对象并关闭这个基础连接，释放资源
	if sqlDB, err := dbBase.DB(); err == nil {
		sqlDB.Close()
	}

	// 3. 构建完整的带有数据库名的 DSN，进行正式连接
	dsnFull := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		APPConfig.DBUser, APPConfig.DBPassword, APPConfig.DBHost, APPConfig.DBPort, APPConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsnFull))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	log.Println("Database connection established")

	// 自动迁移
	db.AutoMigrate(
		&model.Item{},
		&model.ProcessedOrder{},
	)
	log.Println("Database migrated successfully")
	return db, nil
}

func CloseDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}
