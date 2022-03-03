package api

import (
	"GatewayCombat/service/api/base"
	"GatewayCombat/service/api/user"
)

func RegisterRouter() {
	base.Router()
	user.Router()
}
