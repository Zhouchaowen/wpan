package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wpan/errors"
)

func ToResponse(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{
			"code": 0,
			"msg":  "success",
		}
	} else {
		data = gin.H{
			"code": 0,
			"msg":  "success",
			"data": data,
		}
	}
	ctx.JSON(http.StatusOK, data)
}

func ToErrorResponse(ctx *gin.Context, err *errors.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	ctx.JSON(err.StatusCode(), response)
}
