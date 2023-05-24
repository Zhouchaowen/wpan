package handler

import (
	"github.com/gin-gonic/gin"
	"wpan/errors"
	"wpan/model"
	"wpan/utils"
)

type RegisterInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var ri RegisterInfo
	if err := c.BindJSON(&ri); err != nil {
		utils.ToErrorResponse(c, errors.InvalidParams)
		return
	}

	if ri.Username == "" && ri.Password == "" {
		utils.ToErrorResponse(c, errors.InvalidParams)
		return
	}

	user := model.CreateUser(ri.Username, ri.Password, "")
	utils.ToResponse(c, user)
}

func Login(c *gin.Context) {
	var ri RegisterInfo
	if err := c.BindJSON(&ri); err != nil {
		utils.ToErrorResponse(c, errors.InvalidParams)
		return
	}

	user := model.QueryUserInfoByName(ri.Username)
	if user.PassWord != ri.Password {
		utils.ToErrorResponse(c, errors.UnauthorizedAuthFailed)
		return
	}
	tokenString, _ := utils.GenToken(user.Id, user.UserName)

	utils.ToResponse(c, gin.H{"token": tokenString})
}
