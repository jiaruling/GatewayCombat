package api

import (
	"GatewayCombat/service/api/admin"
	"GatewayCombat/service/api/base"
	"GatewayCombat/service/api/user"
)

func RegisterRouter() {
	admin.Router()
	base.Router()
	user.Router()
}
