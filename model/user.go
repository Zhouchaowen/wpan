package model

import (
	"github.com/google/uuid"
	"time"
	"wpan/db"
)

// User 用户表
type User struct {
	Id           string    `json:"id"`
	UserName     string    `json:"user_name"`
	PassWord     string    `json:"pass_word"`
	ImagePath    string    `json:"image_path"`
	RegisterTime time.Time `json:"register_time"`
}

// CreateUser 创建用户并新建文件仓库
func CreateUser(username, password, image string) User {
	db := db.DB

	user := User{
		Id:           uuid.New().String(),
		UserName:     username,
		PassWord:     password,
		RegisterTime: time.Now(),
		ImagePath:    image,
	}
	db.Create(&user)

	fileStore := FileStore{
		Id:          uuid.New().String(),
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048576000,
	}
	db.Create(&fileStore)

	db.Save(&user)

	return user
}

// QueryUserInfoById 根据Id查询用户
func QueryUserInfoById(userId string) (user User) {
	db := db.DB
	db.Find(&user, "id = ?", userId)
	return
}

// QueryUserInfoByName 根据username查询用户
func QueryUserInfoByName(username string) (user User) {
	db := db.DB

	db.Find(&user, "user_name = ?", username)
	return
}
