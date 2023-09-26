package errno

import (
	"errors"
	"fmt"
)

// Errno 定义了 miniblog 使用的错误类型.
type Errno struct {
	HTTP    int32
	Code    string
	Message string
}

//func (e Errno) Error() string {
//	return fmt.Sprintf("http_status_code=%d, err_msg=%s", e.HTTP, e.Message)
//}

// Error 返回了错误信息.
func (err Errno) Error() string {
	return fmt.Sprintf("http_status_code=%d, err_msg=%s", err.HTTP, err.Message)
}

// SetMessage 设置错误信息.
// Example: errno.ErrInvalidParameter.SetMessage("username is empty")
func (err *Errno) SetMessage(format string, args ...interface{}) *Errno {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// Decode 根据参数 err 解析其中的错误信息.
// Example: httpCode, errCode, errMsg := errno.Decode(err)
func Decode(err error) (int32, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:
	}

	// 默认返回未知错误码和错误信息. 该错误代表服务端出错
	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}

// MatchErr 用于识别是否有错误类型是被定义的
func MatchErr(error error) *Errno {
	var err *Errno
	if errors.As(error, &err) {
		return err
	}
	return NewErr(400, "UnknownError", error.Error())
}

func NewErr(httpCode int32, code, msg string) *Errno {
	return &Errno{
		HTTP:    httpCode,
		Code:    code,
		Message: msg,
	}
}
