package repository

import (
	"user-center/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	initMySQL()
}

func initSQLite() {
	initDB, err := gorm.Open(sqlite.Open("user-center"))
	if err != nil {
		panic("init repository gorm err " + err.Error())
	}

	if initDB.Error != nil {
		panic("init repository gorm err " + initDB.Error.Error())
	}

	err = initDB.AutoMigrate(model.User{})
	if err != nil {
		panic(err)
	}

	db = initDB
}

func initMySQL() {
	dns := "root:123456@tcp(127.0.0.1:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"
	initDB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("init repository grom err" + err.Error())
	}
	if initDB.Error != nil {
		panic("init repositroy grom err" + initDB.Error.Error())
	}

	err = initDB.AutoMigrate(model.User{})

	if err != nil {
		panic(err)
	}

	db = initDB
}
