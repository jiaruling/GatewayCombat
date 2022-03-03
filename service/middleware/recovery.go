package middleware

import (
	"GatewayCombat/global/errInfo"
	"GatewayCombat/service/grf"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 捕获所有panic，并且返回错误信息
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/3 14:30
*/

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("【内部错误】" + fmt.Sprint(err))
				grf.Handler500(c, errInfo.InternalUnknownError, err)
				return
			}
		}()
		c.Next()
	}
}
