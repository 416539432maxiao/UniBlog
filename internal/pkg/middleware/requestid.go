package middleware

import (
	"UniBlog/internal/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//RequestID是一个Gin中间件，在每一个HTTP请求的context，response中注入`X-Request-ID`键值对；

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		//检查请求头中是否有`X-Request-ID`,如果有就复用，没有则新建
		RequestID := c.Request.Header.Get(known.XRequestIDKey)

		if RequestID == "" {
			RequestID = uuid.New().String()
		}
		//将RequestID保存在gin.Context中，方便后续程序使用；
		c.Set(known.XRequestIDKey, RequestID)
		//将RequestID报保存在HTTP返回头中，
		c.Writer.Header().Set(known.XRequestIDKey, RequestID)
		c.Next()
	}
}
