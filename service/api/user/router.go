package user

import (
	"GatewayCombat/global"
	"GatewayCombat/service/api/user/controller"
)

func Router() {
	global.GinRouter.V1.Any("/stu/*id", controller.Stus)
}
