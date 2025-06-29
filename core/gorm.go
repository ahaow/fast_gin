package core

import (
	"fast_gin/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(config *config.Config) *gorm.DB {

	database := config.Database

	fmt.Println("database.Mode", database.Mode)

	if database.Mode != "mysql" {
		log.Fatal("未配置数据库")
	}

	dsn := database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成实体外键
	})

	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	sqlDb, err := db.DB()

	if err != nil {
		log.Fatalf("获取数据库连接失败: %v", err)
	}

	err = sqlDb.Ping()

	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	sqlDb.SetMaxIdleConns(database.MaxIdleConns) //最大连接池数量
	sqlDb.SetMaxOpenConns(database.MaxOpenCons)  // 打开数据库最大数量
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("数据库设置配置失败: %v", err)
	}

	return db
}
