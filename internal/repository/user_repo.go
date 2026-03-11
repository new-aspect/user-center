package repository

import (
	"fmt"
	"user-center/internal/model"

	"gorm.io/gorm"
)

func CreateUser(user model.User) (int, error) {
	if db == nil {
		return 0, fmt.Errorf("数据库初始化为空")
	}

	tx := db.Create(&user)
	if tx == nil {
		return 0, fmt.Errorf("db.Create 为空")
	}

	//err := db.Create(user).Error
	//return user.ID, tx.Error
	return 0, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.Where("Username = ?", username).First(&user).Error
	// ignore not found error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, err
}

//func GetUserByID(userID int) (*model.User, error) {
//	var user model.User
//	err := db.Find("ID", userID).First(&user).Error
//	return &user, err
//}
