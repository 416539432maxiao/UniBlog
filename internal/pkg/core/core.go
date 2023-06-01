package core

import (
	"UniBlog/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

//该包用于统一返回方法

/*

要在服务器端返回这样的错误信息给客户端；

{
"code": "InvalidParameter.BadAuthenticationData",
"message": "Bad Authentication data."
}


*/

//定义发生错误时的返回消息；

type ErrResponse struct {
	Code string `json:"code"`

	Message string `json:"message"`
}

//将错误或响应数据写入HTTP的响应主体；

// 使用errno.Decode方法，根据错误类型，尝试从err中提取业务错误码和错误信息；
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
