package middleware

import (
	"GatewayCombat/global"
	"time"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 访问日志
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 17:16
*/

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var bodyData []byte
		//if c.Request.URL.Path != "/api/v1/from/submit/upload" {
		//	// 获取请求体参数
		//	bodyData, _ = c.GetRawData()
		//	// 回写请求体参数
		//	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))
		//}
		defer func(start time.Time) {
			global.AccessLog.Printf("%v || %v || %v || %v ms", c.Request.Method, c.Request.RequestURI, c.Request.RemoteAddr, time.Since(start).Milliseconds())
		}(time.Now())

		// 处理请求
		c.Next()
	}
}
