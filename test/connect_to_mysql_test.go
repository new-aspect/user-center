package test

import (
	"fmt"
	"testing"
	"user-center/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnectMysql(t *testing.T) {
	dns := "root:123456@tcp(127.0.0.1:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println("success")

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	err = db.Create(&model.User{Username: "hello", Password: "123456"}).Error
	if err != nil {
		panic(err)
	}

	var user model.User
	if err = db.First(&user).Error; err != nil {
		panic(err)
	}
	fmt.Println(user)
}
