package config

import (
	"group_project/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/group_project?charset=utf8mb4&parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	MigrateDB(db)
	return db
}

func MigrateDB(conn *gorm.DB) {
	conn.AutoMigrate(entity.Rent{})
	conn.AutoMigrate(entity.Book{})
	conn.AutoMigrate(entity.User{})
}
