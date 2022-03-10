package api

import (
	"GatewayCombat/service/api/admin"
	"GatewayCombat/service/api/app"
	"GatewayCombat/service/api/base"
	"GatewayCombat/service/api/dashboard"
	"GatewayCombat/service/api/service"
)

func RegisterRouter() {
	base.Router()
	admin.Router()
	service.Router()
	app.Router()
	dashboard.Router()
	//user.Router()
}
