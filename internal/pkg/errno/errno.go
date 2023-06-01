package errno

import (
	"fmt"
)

//错误包

// 定义Uniblog使用的错误类型
type Errno struct {
	HTTP    int
	Code    string
	Message string
}

// 实现error接口的方法；
func (err *Errno) Error() string {
	return err.Message
}

// 设置好Errno结构体实例的Message字段
func (err *Errno) SetMessage(format string, args ...interface{}) *Errno {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// Decode尝试从err中解析出业务错误码和错误信息
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}
	//判断如果是*Errno:的数据类型，返回相应的错误码和错误信息
	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:
	}

	// 默认返回未知错误码和错误信息. 该错误代表服务端出错
	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}
