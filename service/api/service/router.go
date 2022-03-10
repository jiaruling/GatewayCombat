package service

import (
	"GatewayCombat/global"
	"GatewayCombat/service/api/service/controller"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 14:07
*/
func Router() {
	r := global.GinRouter.V1.Group("/service")
	controller.ServiceRegister(r)
}
