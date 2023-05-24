package model

import (
	"github.com/google/uuid"
	"time"
	"wpan/db"
)

// Folder 文件夹表
type Folder struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	UserId    string    `json:"user_id"`
	ParentId  string    `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateFolder 新建文件夹
func CreateFolder(userId, name, parentId string) (Folder, error) {
	db := db.DB
	folder := Folder{
		Id:       uuid.New().String(),
		UserId:   userId,
		Name:     name,
		ParentId: parentId,
	}
	err := db.Create(&folder).Error

	return folder, err
}

// QueryFolderById 获取当前的目录信息
func QueryFolderById(userId, folderId string) (folder Folder) {
	db := db.DB
	db.Find(&folder, "user_id = ? and id = ?", userId, folderId)
	return
}

// QueryFolders 获取目录所有文件夹
func QueryFolders(usrId, folderId string) (folders []Folder) {
	db := db.DB
	db.Order("updated_at desc").Find(&folders, "user_id = ? and parent_id = ?", usrId, folderId)
	return
}

// UpdateFolderName 修改文件夹名
func UpdateFolderName(userId, folderId, name string) (Folder, error) {
	db := db.DB
	var folder Folder
	err := db.Model(&folder).Where("user_id = ? and id = ?", userId, folderId).Update("name", name).Error
	return folder, err
}

// DeleteFolder 删除文件夹信息
func DeleteFolder(userId, folderId string) {
	db := db.DB

	//删除文件夹信息
	db.Where("user_id = ? and id = ?", userId, folderId).Delete(Folder{})
	//删除文件夹中文件夹信息
	db.Where("user_id = ? and parent_id = ?", userId, folderId).Delete(Folder{})
	//删除文件夹中文件信息
	db.Where("user_id = ? and folder_id = ?", userId, folderId).Delete(File{})
}
