package model

import (
	"fmt"
	"github.com/google/uuid"
	"path"
	"strings"
	"time"
	"wpan/db"
	"wpan/utils"
)

// File 文件表
type File struct {
	Id          string    `json:"id"`
	UserId      string    `json:"user_id"`
	FolderId    string    `json:"folder_id"`                                                                  //父文件夹id
	FileName    string    `json:"file_name" sql:"type:VARCHAR(5) CHARACTER SET utf8 COLLATE utf8_general_ci"` //文件名
	FileHash    string    `json:"file_hash"`                                                                  //文件哈希值
	Postfix     string    `json:"postfix"`                                                                    //文件后缀
	Size        int64     `json:"size"`                                                                       //文件大小
	Type        int       `json:"type"`                                                                       //文件类型
	DownloadNum int       `json:"download_num"`                                                               //下载次数
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateFile 添加文件数据
func CreateFile(filename, fileHash, folderId, userId string, fileSize int64) (File, error) {
	db := db.DB
	//获取文件后缀
	fileSuffix := path.Ext(filename)
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	file := File{
		Id:          uuid.New().String(),
		UserId:      userId,
		FolderId:    folderId,
		FileName:    filePrefix,
		FileHash:    fileHash,
		Postfix:     strings.ToLower(fileSuffix),
		Size:        fileSize / 1024,
		Type:        utils.GetFileTypeInt(fileSuffix),
		DownloadNum: 0,
	}
	err := db.Create(&file).Error
	return file, err
}

// SubtractSize 文件上传成功减去相应容量
func SubtractSize(size int64, userId string) error {
	db := db.DB
	var fileStore FileStore
	db.First(&fileStore, "user_id = ?", userId)

	fileStore.CurrentSize = fileStore.CurrentSize + size
	fileStore.MaxSize = fileStore.MaxSize - size
	return db.Save(&fileStore).Error
}

// QueryUserFile 获取用户指定目录下的文件
func QueryUserFile(userId, folderId string) (files []File) {
	db := db.DB
	db.Find(&files, "user_id = ? and folder_id = ?", userId, folderId)
	return
}

// QueryTypeFile 根据文件类型获取文件
func QueryTypeFile(fileType int, userId string) (files []File) {
	db := db.DB
	db.Find(&files, "type = ? and user_id = ?", fileType, userId)
	return
}

// QueryFileInfoById 通过fileId获取文件信息
func QueryFileInfoById(fileId string) (file File) {
	db := db.DB
	db.First(&file, "id = ?", fileId)
	return
}

// DownloadNumAdd 文件下载次数+1
func DownloadNumAdd(fileId string) error {
	var file File
	db := db.DB

	db.First(&file, "id = ?", fileId)
	if len(file.Id) == 0 {
		return fmt.Errorf("no file data")
	}

	file.DownloadNum = file.DownloadNum + 1
	return db.Save(&file).Error
}

// DeleteUserFile 删除数据库文件数据
func DeleteUserFile(userId, fileId string) error {
	db := db.DB
	return db.Delete(&File{}, "user_id = ? and id = ?", userId, fileId).Error
}

// QueryFileDetailUse 获取用户文件使用明细情况
func QueryFileDetailUse(userId string) map[string]int64 {
	var files []File
	var (
		docCount   int64
		imgCount   int64
		videoCount int64
		musicCount int64
		otherCount int64
	)
	db := db.DB

	fileDetailUseMap := make(map[string]int64, 0)

	//文档类型
	docCount = db.Find(&files, "user_id = ? and type = ?", userId, 1).RowsAffected
	fileDetailUseMap["docCount"] = docCount
	////图片类型
	imgCount = db.Find(&files, "user_id = ? and type = ?", userId, 2).RowsAffected
	fileDetailUseMap["imgCount"] = imgCount
	//视频类型
	videoCount = db.Find(&files, "user_id = ? and type = ?", userId, 3).RowsAffected
	fileDetailUseMap["videoCount"] = videoCount
	//音乐类型
	musicCount = db.Find(&files, "user_id = ? and type = ?", userId, 4).RowsAffected
	fileDetailUseMap["musicCount"] = musicCount
	//其他类型
	otherCount = db.Find(&files, "user_id = ? and type = ?", userId, 5).RowsAffected
	fileDetailUseMap["otherCount"] = otherCount

	return fileDetailUseMap
}
