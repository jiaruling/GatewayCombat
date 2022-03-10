package dto

import "time"

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/4 15:50
*/

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	Password string `json:"password" dto:"password" comment:"密码" example:"123456" validate:"required"` //密码
}
