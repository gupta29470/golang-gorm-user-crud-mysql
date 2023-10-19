package databases

import (
	"github.com/gupta29470/golang-sql-crud-with-orm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

func ConnectDB() {
	var gormOpenError error
	dsn := "user_name:password@tcp(localhost:3306)/go_user_sql_orm?charset=utf8mb4&parseTime=True&loc=Local"
	dB, gormOpenError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if gormOpenError != nil {
		panic("Failed to connect database")
	}
}

func InitDB() {
	ConnectDB()
	dB.AutoMigrate(&models.User{})
}

func DB() *gorm.DB {
	return dB
}
