package handler

import (
	"fmt"
	"net/http"
	"user-center/internal/model"
	"user-center/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostUserRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		Fail(c, http.StatusBadRequest, "输入格式不对："+err.Error())
		return
	}

	existUser, err := repository.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if existUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
		return
	}

	user.Password, err = encryptedPasswordUsingBcrypt(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userInDB, err := repository.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userInDB})
}

func encryptedPasswordUsingBcrypt(password string) (string, error) {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(generateFromPassword), err
}

func PostUserLogin(c *gin.Context) {
	var loginUser model.User

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := repository.GetUserByUsername(loginUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if dbUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		return
	}

	isOk := verifyPasswordUsingBcryptIsOk(dbUser.Password, loginUser.Password)
	if !isOk {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": dbUser})
}

func verifyPasswordUsingBcryptIsOk(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
