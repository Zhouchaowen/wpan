package handler

import (
	"github.com/gin-gonic/gin"
	"wpan/errors"
	"wpan/middleware"
	"wpan/model"
	"wpan/utils"
)

type CreateFolderInfo struct {
	ParentId   string `json:"parent_id"`
	FolderName string `json:"folder_name"`
}

// CreateFolder 处理新建文件夹
func CreateFolder(c *gin.Context) {
	var cfi CreateFolderInfo
	if err := c.BindJSON(&cfi); err != nil {
		utils.ToErrorResponse(c, errors.InvalidParams)
		return
	}
	//获取用户Id
	userId := c.GetString(middleware.UserIdKey)

	//新建文件夹数据
	child, err := model.CreateFolder(userId, cfi.FolderName, cfi.ParentId)
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.FolderCreateFailed)
		return
	}

	utils.ToResponse(c, child)
}

type UpdateFolderInfo struct {
	FolderId   string `json:"folder_id"`
	FolderName string `json:"folder_name"`
}

// UpdateFolder 处理新建文件夹
func UpdateFolder(c *gin.Context) {
	var ufi UpdateFolderInfo
	if err := c.BindJSON(&ufi); err != nil {
		utils.ToErrorResponse(c, errors.InvalidParams)
	}
	//获取用户Id
	userId := c.GetString(middleware.UserIdKey)

	//新建文件夹数据
	child, err := model.UpdateFolderName(userId, ufi.FolderId, ufi.FolderName)
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.FolderUpdateFailed)
		return
	}

	utils.ToResponse(c, child)
}

// Folders 获取目录下所有文件，文件夹
func Folders(c *gin.Context) {
	userId := c.GetString(middleware.UserIdKey)
	folderId := c.Param("folder_id")
	//获取当前目录所有文件
	files := model.QueryUserFile(userId, folderId)

	//获取当前目录所有文件夹
	folders := model.QueryFolders(userId, folderId)

	//获取当前目录信息
	currentFolder := model.QueryFolderById(userId, folderId)

	if len(currentFolder.Id) == 0 {
		utils.ToResponse(c, gin.H{
			"files":          files,
			"folders":        folders,
			"current_folder": nil,
		})
		return
	}

	utils.ToResponse(c, gin.H{
		"files":          files,
		"folders":        folders,
		"current_folder": currentFolder,
	})
}

// DeleteFolder 删除文件夹
func DeleteFolder(c *gin.Context) {
	userId := c.GetString(middleware.UserIdKey)
	folderId := c.Param("folder_id")

	model.DeleteFolder(userId, folderId)

	utils.ToResponse(c, "")
}
