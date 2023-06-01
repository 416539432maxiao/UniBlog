package middleware

import "github.com/gin-gonic/gin"

//cors 跨域资源请求中间件用来设置 options 请求的返回头，然后退出中间件链，并结束请求(浏览器跨域设置).

func Cors(c *gin.Context) {
	//如果 HTTP 请求不是 OPTIONS 跨域请求，则继续处理 HTTP 请求；
	if c.Request.Method != "OPTIONS" {
		c.Next()
		//如果 HTTP 请求是OPTIONS 跨域请求，则设置跨域 Header，并返回。
	} else {
		//对跨域请求的权限进行设置；
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//NoCache中间件，用来禁止客户端缓存HTTP请求的返回结果；
