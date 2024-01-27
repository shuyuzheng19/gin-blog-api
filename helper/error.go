package helper

import (
	"common-web-framework/common"
	"errors"
)

// ErrorPanicAndMessage 全局error处理
func ErrorPanicAndMessage(err error, message string) {
	if err != nil {
		var err interface{} = errors.New(message + "||" + err.Error())
		panic(err)
	}
}

// ErrorToResponse  全局error处理 返回给前端
func ErrorToResponse(errCode common.ErrorCode) {
	var r interface{} = common.AutoFail(errCode)
	panic(r)
}
