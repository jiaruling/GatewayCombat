package controller

import (
	"GatewayCombat/global"
	"GatewayCombat/service/api/admin/dao"
	"GatewayCombat/service/api/admin/dto"
	"GatewayCombat/service/grf"
	"encoding/json"
	"time"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 管理员登录和退出
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/4 10:34
*/

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLoginOut)
}

// AdminLogin godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (a *AdminLoginController) AdminLogin(c *gin.Context) {
	// 表单验证
	req := &dto.AdminLoginInput{}
	if err := c.ShouldBindJSON(&req); err != nil {
		grf.FormsVerifyFailed(c, err)
		return
	}
	// 验证用户名密码
	adminDao := &dao.AdminDao{}
	admin, err := adminDao.LoginCheck(global.RDB, req)
	if err != nil {
		grf.Handler400(c, err.Error(), nil)
		return
	}

	// 设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		grf.Handler500(c, "", err)
		return
	}
	sess := sessions.Default(c)
	sess.Set(global.AdminSessionInfoKey, string(sessBts))
	sess.Save()

	out := &dto.AdminLoginOutput{Token: admin.UserName}
	grf.Handler200(c, out)
}

// AdminLogin godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [get]
func (a *AdminLoginController) AdminLoginOut(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(global.AdminSessionInfoKey)
	sess.Save()
	grf.Handler200(c, "")
}
