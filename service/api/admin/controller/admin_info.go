package controller

import (
	"GatewayCombat/global"
	"GatewayCombat/service/api/admin/dao"
	"GatewayCombat/service/api/admin/dto"
	"GatewayCombat/service/grf"
	"GatewayCombat/service/public"
	"GatewayCombat/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/4 10:34
*/
type AdminController struct{}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/info", adminLogin.AdminInfo)
	group.POST("/change/pwd", adminLogin.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (a *AdminController) AdminInfo(c *gin.Context) {
	//1. 读取sessionKey对应json 转换为结构体
	sess := sessions.Default(c)
	sessInfo := sess.Get(global.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		grf.Handler500(c, err.Error(), nil)
		return
	}
	//2. 取出数据然后封装输出结构体
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	grf.Handler200(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (a *AdminController) ChangePwd(c *gin.Context) {
	req := &dto.ChangePwdInput{}
	if err := c.ShouldBind(req); err != nil {
		public.FormsVerifyFailed(c, err)
		return
	}
	//1. session读取用户信息到结构体 sessInfo
	sess := sessions.Default(c)
	sessInfo := sess.Get(global.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		grf.Handler500(c, err.Error(), nil)
		return
	}
	//2. sessInfo.ID 读取数据库信息 adminInfo
	admin := &dao.Admin{}
	adminInfo, err := admin.Find(global.RDB, &dao.Admin{Id: adminSessionInfo.ID})
	if err != nil {
		grf.Handler500(c, err.Error(), nil)
		return
	}
	//3. params.password+adminInfo.salt sha256 saltPassword
	saltPassword := utils.GenSaltPassword(adminInfo.Salt, req.Password)
	adminInfo.Password = saltPassword
	//4. saltPassword==> adminInfo.password 执行数据保存
	if err := admin.Save(global.WDB, adminInfo); err != nil {
		grf.Handler500(c, err.Error(), nil)
		return
	}
	grf.Handler200(c, "")
}
