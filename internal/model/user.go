package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `binding:"required"`
	Password string `binding:"required"`
}
