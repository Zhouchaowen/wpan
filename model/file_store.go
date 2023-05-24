package model

import "wpan/db"

type FileStore struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	CurrentSize int64  `json:"current_size"`
	MaxSize     int64  `json:"max_size"`
}

// CapacityIsEnough 判断用户容量是否足够
func CapacityIsEnough(fileSize int64, UserId string) bool {
	db := db.DB
	var fileStore FileStore
	db.First(&fileStore, "user_id = ?", UserId)
	if fileStore.MaxSize-(fileSize/1024) < 0 {
		return false
	}
	return true
}
