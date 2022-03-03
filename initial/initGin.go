package initial

import (
	"GatewayCombat/global"
	m "GatewayCombat/service/middleware"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 添加全局中间件
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 15:42
*/

func InitGin() {
	// 初始gin的路由并赋值给全局变量
	r := gin.Default()
	// 注册全局中间件，跨域请求
	r.Use(m.RecoveryMiddleware(), m.Cors(), m.AccessLog())
	apiV1 := r.Group("/api/v1")
	// 复制给全局单例
	global.GinRouter = &global.Router{
		Router: r,
		V1:     apiV1,
	}

	// 控制台显示日志显示颜色
	gin.ForceConsoleColor()

	// 设定gin服务器启动的模式
	runMode := global.Config.RunMode
	if runMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
