package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
	"strings"
	"wpan/config"
	"wpan/errors"
	"wpan/middleware"
	"wpan/model"
	"wpan/utils"
)

// CreateShare 创建分享
func CreateShare(c *gin.Context) {
	//获取用户Id
	userId := c.GetString(middleware.UserIdKey)
	fileId := c.Param("file_id")

	code := strings.ReplaceAll(uuid.New().String(), "-", "")[:10]
	//新建分享
	share, err := model.CreateShare(code, userId, fileId)
	if err != nil {
		c.Error(err)
		utils.ToErrorResponse(c, errors.ShareCreateFailed)
		return
	}

	utils.ToResponse(c, share)
}

func QueryShare(c *gin.Context) {
	id := c.Param("id")
	code := c.Query("code")

	share := model.QueryShare(id)
	if share.Code != code {
		utils.ToErrorResponse(c, errors.FileNotFound)
		return
	}

	file := model.QueryFileInfoById(share.FileId)
	utils.ToResponse(c, file)
}

func QueryShares(c *gin.Context) {
	//获取用户Id
	userId := c.GetString(middleware.UserIdKey)
	shares := model.QuerySharesByUserId(userId)
	utils.ToResponse(c, shares)
}

func ShareDownloadFile(c *gin.Context) {
	id := c.Param("id")
	code := c.Query("code")
	share := model.QueryShare(id)
	if share.Code != code {
		utils.ToErrorResponse(c, errors.FileNotFound)
		return
	}

	file := model.QueryFileInfoById(share.FileId)
	if file.FileHash == "" {
		utils.ToErrorResponse(c, errors.FileNotFound)
		return
	}

	location := path.Join(config.ServerConfig.APP.Location, file.FileName+file.Postfix)
	fmt.Println(location)
	//下载次数+1
	model.DownloadNumAdd(share.FileId)

	c.File(location)
}

func DeleteShare(c *gin.Context) {
	//获取用户Id
	userId := c.GetString(middleware.UserIdKey)
	id := c.Param("id")

	err := model.DeleteShare(id, userId)
	if err != nil {
		c.Error(err)
		utils.ToResponse(c, errors.FileDeleteFailed)
		return
	}
	utils.ToResponse(c, "")
}
