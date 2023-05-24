package errors

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "服务内部错误")
	InvalidParams             = NewError(10001, "入参错误")
	NotFound                  = NewError(10002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10003, "账户不存在")
	UnauthorizedAuthFailed    = NewError(10004, "账户密码错误")
	UnauthorizedTokenError    = NewError(10005, "鉴权失败，Token 错误或丢失")
	UnauthorizedTokenTimeout  = NewError(10006, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10007, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10008, "请求过多")

	GatewayMethodsLimit    = NewError(10109, "网关仅接受GET或POST请求")
	GatewayLostSign        = NewError(10110, "网关请求缺少签名")
	GatewayLostAppKey      = NewError(10111, "网关请求缺少APP KEY")
	GatewayAppKeyInvalid   = NewError(10112, "网关请求无效APP KEY")
	GatewayAppKeyClosed    = NewError(10113, "网关请求APP KEY已停用")
	GatewayParamSignError  = NewError(10114, "网关请求参数签名错误")
	GatewayTooManyRequests = NewError(10115, "网关请求频次超限")

	FileUploadFailed  = NewError(10200, "文件上传失败")
	FileInvalidExt    = NewError(10201, "文件类型不合法")
	FileInvalidSize   = NewError(10202, "文件大小超限")
	FileUnderCapacity = NewError(10203, "容量不足")
	FileNotFound      = NewError(10204, "文件不存在")
	FileDeleteFailed  = NewError(10205, "文件删除失败")

	FolderCreateFailed = NewError(20000, "文件夹创建失败")
	FolderUpdateFailed = NewError(20001, "文件夹更新失败")

	ShareCreateFailed = NewError(30000, "分享创建失败")
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedAuthFailed.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
