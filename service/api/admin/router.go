package admin

import (
	"GatewayCombat/global"
	"GatewayCombat/service/api/admin/controller"
)

/*
   功能说明: 管理员接口
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/4 10:30
*/

func Router() {
	r := global.GinRouter.V1.Group("/admin")
	controller.AdminLoginRegister(r)
	controller.AdminRegister(r)
}
