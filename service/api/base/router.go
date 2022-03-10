package base

import (
	"GatewayCombat/global"
	. "GatewayCombat/service/api/base/controller"
	"net/http"
)

func Router() {
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	global.GinRouter.Router.StaticFS("/static", http.Dir(global.FILEPATH))
	global.GinRouter.Router.StaticFS("/log", http.Dir(global.LogFilePath))
	global.GinRouter.Router.NoRoute(NotFound)      // 404 路由
	global.GinRouter.Router.GET("/health", Health) // 服务健康检查
	global.GinRouter.Router.GET("/info", Info)     // info
	global.GinRouter.Router.GET("/ping", Ping)     // Ping
	global.GinRouter.Router.GET("/file/base64", GetConfigFile)
	global.GinRouter.Router.POST("/file/base64", PostConfigFile)
}
