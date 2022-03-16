package controller

import "github.com/gin-gonic/gin"

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/3/10 17:03
*/

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	dashboard := &DashboardController{}
	group.GET("/panel_group_data", dashboard.PanelGroupData)
	group.GET("/flow_stat", dashboard.FlowStat)
	group.GET("/service_stat", dashboard.ServiceStat)
}

func (dc *DashboardController) PanelGroupData(ctx *gin.Context) {

}

func (dc *DashboardController) FlowStat(ctx *gin.Context) {

}

func (dc *DashboardController) ServiceStat(ctx *gin.Context) {

}
