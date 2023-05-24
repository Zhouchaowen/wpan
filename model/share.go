package model

import (
	"github.com/google/uuid"
	"time"
	"wpan/db"
)

type Share struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	UserId    string    `json:"user_id"`
	FileId    string    `json:"file_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateShare 添加分享数据
func CreateShare(code, userId, fileId string) (Share, error) {
	db := db.DB

	share := Share{
		Id:     uuid.New().String(),
		Code:   code,
		UserId: userId,
		FileId: fileId,
	}
	err := db.Create(&share).Error
	return share, err
}

// QueryShare 获取分享
func QueryShare(id string) (share Share) {
	db := db.DB
	db.Find(&share, "id = ?", id)
	return
}

// QuerySharesByUserId 获取分享
func QuerySharesByUserId(userId string) (shares []Share) {
	db := db.DB
	db.Find(&shares, "user_id = ?", userId)
	return
}

// DeleteShare 删除数据库文件数据
func DeleteShare(id, userId string) error {
	db := db.DB
	return db.Delete(&File{}, "id = ? and user_id = ?", id, userId).Error
}
