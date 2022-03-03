package initial

import (
	"GatewayCombat/global"
	"fmt"
	"log"
)

/*
   功能说明: 启动服务
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 14:45
*/

func InitService() {
	go func() {
		if err := global.GinRouter.Router.Run(fmt.Sprintf("0.0.0.0:%v", 8080)); err != nil {
			log.Fatalln("http服务启动失败: ", err.Error())
		}
	}()
}
