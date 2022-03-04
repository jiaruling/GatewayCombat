package api

import (
	"GatewayCombat/service/api/admin"
	"GatewayCombat/service/api/base"
)

func RegisterRouter() {
	admin.Router()
	base.Router()
	//user.Router()
}
