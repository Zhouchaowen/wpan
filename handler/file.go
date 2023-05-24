package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"wpan/config"
	"wpan/errors"
	"wpan/middleware"
	"wpan/model"
	"wpan/utils"
)

// UploadFile 处理上传文件
func UploadFile(c *gin.Context) {
	//获取用户Id
	userId := c.GetString(middleware.UserIdKey)
	folderId, ok := c.GetPostForm("folder_id")
	if !ok {
		utils.ToErrorResponse(c, errors.InvalidParams)
		return
	}

	conf := config.ServerConfig
	//接收上传文件
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.FileUploadFailed)
		return
	}
	defer file.Close()

	//判断用户的容量是否足够
	if ok := model.CapacityIsEnough(head.Size, userId); !ok {
		utils.ToErrorResponse(c, errors.FileUnderCapacity)
		return
	}

	//文件保存本地的路径
	location := path.Join(conf.APP.Location, head.Filename)

	//在本地创建一个新的文件
	newFile, err := os.Create(location)
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.FileUploadFailed)
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.FileUploadFailed)
		return
	}

	//将光标移至开头
	_, _ = newFile.Seek(0, 0)
	fileHash := utils.GetSHA256HashCode(newFile)

	//新建文件信息
	ret, err := model.CreateFile(head.Filename, fileHash, folderId, userId, fileSize)
	if err != nil {
		c.Error(err)
		utils.ToResponse(c, errors.FileUploadFailed)
		return
	}
	//上传成功减去相应剩余容量
	model.SubtractSize(fileSize, userId)

	utils.ToResponse(c, ret)
}

func QueryFile(c *gin.Context) {
	fileId := c.Param("file_id")
	file := model.QueryFileInfoById(fileId)
	if file.FileHash == "" {
		utils.ToErrorResponse(c, errors.FileNotFound)
		return
	}
	utils.ToResponse(c, file)
}

// DownloadFile 下载文件
func DownloadFile(c *gin.Context) {
	fileId := c.Param("file_id")

	file := model.QueryFileInfoById(fileId)
	if file.FileHash == "" {
		utils.ToErrorResponse(c, errors.FileNotFound)
		return
	}

	location := path.Join(config.ServerConfig.APP.Location, file.FileName+file.Postfix)
	fmt.Println(location)
	//下载次数+1
	model.DownloadNumAdd(fileId)

	c.File(location)
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	userId := c.GetString(middleware.UserIdKey)
	fileId := c.Param("file_id")

	//删除数据库文件数据
	err := model.DeleteUserFile(userId, fileId)
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.FileDeleteFailed)
	}

	utils.ToResponse(c, "")
}

func TypeFiles(c *gin.Context) {
	userId := c.GetString(middleware.UserIdKey)
	typeNum := c.GetInt("type_id")
	//获取文档类型文件
	typeFiles := model.QueryTypeFile(typeNum, userId)

	utils.ToResponse(c, gin.H{
		"type_files": typeFiles,
		"type_count": len(typeFiles),
	})
}
